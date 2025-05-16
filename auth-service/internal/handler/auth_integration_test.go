package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

const baseURL = "http://localhost:8081/api/v1"

func TestRegisterLoginProfile(t *testing.T) {
	// 1. Register
	registerBody := map[string]string{
		"email":      "testuser1@example.com",
		"password":   "Test123!@#",
		"first_name": "Test",
		"last_name":  "User",
		"role":       "student",
	}
	regData, _ := json.Marshal(registerBody)
	resp, err := http.Post(baseURL+"/auth/register", "application/json", bytes.NewReader(regData))
	if err != nil {
		t.Fatalf("register request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("register failed: %s", string(b))
	}

	// 2. Login
	loginBody := map[string]string{
		"email":    "testuser1@example.com",
		"password": "Test123!@#",
	}
	loginData, _ := json.Marshal(loginBody)
	resp, err = http.Post(baseURL+"/auth/login", "application/json", bytes.NewReader(loginData))
	if err != nil {
		t.Fatalf("login request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("login failed: %s", string(b))
	}
	var loginResp struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		t.Fatalf("decode login response: %v", err)
	}
	if loginResp.AccessToken == "" {
		t.Fatal("no access_token in login response")
	}

	// 3. Get Profile
	req, _ := http.NewRequest("GET", baseURL+"/users/profile", nil)
	req.Header.Set("Authorization", "Bearer "+loginResp.AccessToken)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("profile request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("profile failed: %s", string(b))
	}
	// Можно добавить проверку структуры профиля
}

func TestRefreshAndChangePassword(t *testing.T) {
	// Register + Login
	registerBody := map[string]string{
		"email":      "testuser2@example.com",
		"password":   "Test123!@#",
		"first_name": "Test",
		"last_name":  "User",
		"role":       "student",
	}
	regData, _ := json.Marshal(registerBody)
	http.Post(baseURL+"/auth/register", "application/json", bytes.NewReader(regData))

	loginBody := map[string]string{
		"email":    "testuser2@example.com",
		"password": "Test123!@#",
	}
	loginData, _ := json.Marshal(loginBody)
	resp, err := http.Post(baseURL+"/auth/login", "application/json", bytes.NewReader(loginData))
	if err != nil {
		t.Fatalf("login request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("login failed: %s", string(b))
	}
	var loginResp struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		t.Fatalf("decode login response: %v", err)
	}

	// 1. Refresh token
	refreshBody := map[string]string{"refresh_token": loginResp.RefreshToken}
	refreshData, _ := json.Marshal(refreshBody)
	resp, err = http.Post(baseURL+"/auth/refresh", "application/json", bytes.NewReader(refreshData))
	if err != nil {
		t.Fatalf("refresh request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("refresh failed: %s", string(b))
	}
	var refreshResp struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&refreshResp); err != nil {
		t.Fatalf("decode refresh response: %v", err)
	}
	if refreshResp.AccessToken == "" {
		t.Fatal("no access_token in refresh response")
	}

	// 2. Change password
	changeBody := map[string]string{
		"old_password": "Test123!@#",
		"new_password": "NewPass123!@#",
	}
	changeData, _ := json.Marshal(changeBody)
	req, _ := http.NewRequest("POST", baseURL+"/users/change-password", bytes.NewReader(changeData))
	req.Header.Set("Authorization", "Bearer "+refreshResp.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("change password request failed: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("change password failed: %s", string(b))
	}

	// 3. Проверка входа со старым паролем (должен быть отказ)
	loginBody["password"] = "Test123!@#"
	loginData, _ = json.Marshal(loginBody)
	resp, _ = http.Post(baseURL+"/auth/login", "application/json", bytes.NewReader(loginData))
	if resp.StatusCode == 200 {
		t.Fatal("login with old password should fail")
	}

	// 4. Проверка входа с новым паролем (должен быть успех)
	loginBody["password"] = "NewPass123!@#"
	loginData, _ = json.Marshal(loginBody)
	resp, err = http.Post(baseURL+"/auth/login", "application/json", bytes.NewReader(loginData))
	if err != nil {
		t.Fatalf("login with new password failed: %v", err)
	}
	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("login with new password failed: %s", string(b))
	}
}

func TestProfileAccessDenied(t *testing.T) {
	// Без токена
	resp, err := http.Get(baseURL+"/users/profile")
	if err != nil {
		t.Fatalf("profile request (no token) failed: %v", err)
	}
	if resp.StatusCode == 200 {
		t.Fatal("profile access without token should be denied")
	}
	resp.Body.Close()

	// С невалидным токеном
	req, _ := http.NewRequest("GET", baseURL+"/users/profile", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("profile request (invalid token) failed: %v", err)
	}
	if resp.StatusCode == 200 {
		t.Fatal("profile access with invalid token should be denied")
	}
	resp.Body.Close()
}

// Для запуска: go test -v ./internal/handler -run Integration 
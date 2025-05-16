package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/krawwwwy/virtual-decanat/auth-service/internal/models"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/repository"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/services"
)

// AuthHandler обрабатывает запросы, связанные с аутентификацией
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler создает новый экземпляр AuthHandler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// RegisterHandler обрабатывает запрос на регистрацию нового пользователя
func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим тело запроса
	var req models.UserRegistration
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Регистрируем пользователя
	user, err := h.authService.Register(req)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		
		if errors.Is(err, repository.ErrUserAlreadyExists) {
			http.Error(w, "User with this username or email already exists", http.StatusConflict)
			return
		}
		
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Формируем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Преобразуем пользователя в UserResponse
	userResponse := models.UserResponse{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		Role:       user.Role,
		Group:      user.Group,
		Faculty:    user.Faculty,
		Department: user.Department,
		CreatedAt:  user.CreatedAt,
	}

	// Отправляем ответ
	json.NewEncoder(w).Encode(userResponse)
}

// LoginHandler обрабатывает запрос на аутентификацию пользователя
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Парсим тело запроса
	var req models.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Аутентифицируем пользователя
	tokenResponse, err := h.authService.Login(req)
	if err != nil {
		log.Printf("Error logging in: %v", err)
		
		if errors.Is(err, repository.ErrInvalidCredentials) {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		
		http.Error(w, "Failed to login", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResponse)
}

// GetProfileHandler возвращает профиль аутентифицированного пользователя
func (h *AuthHandler) GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из запроса (предполагается, что middleware добавил его)
	userID, err := extractUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Получаем профиль пользователя
	userResponse, err := h.authService.GetUserByID(userID)
	if err != nil {
		log.Printf("Error getting user profile: %v", err)
		
		if errors.Is(err, repository.ErrUserNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		
		http.Error(w, "Failed to get user profile", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}

// UpdateProfileHandler обновляет профиль аутентифицированного пользователя
func (h *AuthHandler) UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из запроса
	userID, err := extractUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Парсим тело запроса
	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Обновляем профиль пользователя
	userResponse, err := h.authService.UpdateUser(userID, updateData)
	if err != nil {
		log.Printf("Error updating user profile: %v", err)
		
		if errors.Is(err, repository.ErrUserNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		
		http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}

// ChangePasswordHandler изменяет пароль аутентифицированного пользователя
func (h *AuthHandler) ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID пользователя из запроса
	userID, err := extractUserIDFromRequest(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Парсим тело запроса
	var req struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Изменяем пароль
	err = h.authService.ChangePassword(userID, req.OldPassword, req.NewPassword)
	if err != nil {
		log.Printf("Error changing password: %v", err)
		
		if errors.Is(err, repository.ErrInvalidCredentials) {
			http.Error(w, "Invalid old password", http.StatusUnauthorized)
			return
		}
		
		http.Error(w, "Failed to change password", http.StatusInternalServerError)
		return
	}

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password changed successfully"})
}

// AuthMiddleware проверяет наличие валидного JWT токена
func (h *AuthHandler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Проверяем формат токена
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
			return
		}

		// Проверяем токен
		claims, err := h.authService.VerifyToken(parts[1])
		if err != nil {
			log.Printf("Error verifying token: %v", err)
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Добавляем информацию о пользователе в контекст запроса
		r.Header.Set("X-User-ID", strconv.Itoa(claims.UserID))
		r.Header.Set("X-User-Role", claims.Role)

		// Передаем запрос следующему обработчику
		next.ServeHTTP(w, r)
	})
}

// Вспомогательная функция для извлечения ID пользователя из запроса
func extractUserIDFromRequest(r *http.Request) (int, error) {
	userIDStr := r.Header.Get("X-User-ID")
	if userIDStr == "" {
		return 0, errors.New("user ID not found in request")
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, err
	}

	return userID, nil
} 
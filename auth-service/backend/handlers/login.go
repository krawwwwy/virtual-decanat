package handlers

import (
	"auth-service/models"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func HandleLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing form:", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		user := models.User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Role:     "-",
		}

		log.Printf("Login request: %+v\n", user)

		authenticated, err := Login(db, user.Username, user.Password)
		if err != nil {
			log.Println("Error authenticating:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if authenticated {
			_, err2 := w.Write([]byte("Login successful!"))
			if err2 != nil {
				fmt.Println("Error writing response:", err2)
				return
			}
		} else {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}
	}
}

func Login(db *sql.DB, username, password string) (bool, error) {
	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
	if err != nil {
		log.Printf("Cant get data to login as %v:", username, err)
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	return err == nil, nil
}

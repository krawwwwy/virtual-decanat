package handlers

import (
	"auth-service/models"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func HandleRegister(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println("Error parsing form: Register page", err)
			http.Error(w, "Invalid request payload. Error parsing form", http.StatusBadRequest)
			return
		}

		user := models.User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Role:     r.FormValue("role"),
		}

		log.Printf("Register request: %+v\n", user)

		err = Register(db, user.Username, user.Password, user.Role)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		_, err = w.Write([]byte("Registration successful!"))
		if err != nil {
			log.Println("Error writing response:", err)
			return
		}
	}
}

func Register(db *sql.DB, username, password, roleName string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error beginning transaction:", err)
		return err
	}

	var userID int
	err = tx.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", username, string(hashedPassword)).Scan(&userID)
	if err != nil {
		log.Println("Error inserting user:", err)
		err := tx.Rollback()
		if err != nil {
			fmt.Println("Error rolling back:", err)
			return err
		}
		return err
	}

	var roleID int
	err = tx.QueryRow("SELECT id FROM roles WHERE name = $1", roleName).Scan(&roleID)
	if err != nil {
		log.Println("Error selecting role:", err)
		err := tx.Rollback()
		if err != nil {
			fmt.Println("Error rolling back:", err)
			return err
		}
		return err
	}

	_, err = tx.Exec("INSERT INTO user_roles (user_id, role_id) VALUES ($1, $2)", userID, roleID)
	if err != nil {
		log.Println("Error inserting user role:", err)
		err := tx.Rollback()
		if err != nil {
			fmt.Println("Error rolling back:", err)
			return err
		}
		return err
	}

	return tx.Commit()
}

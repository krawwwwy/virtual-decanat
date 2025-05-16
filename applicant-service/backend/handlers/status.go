package handlers

import (
	"applicant-service/models"
	"database/sql"
	"html/template"
	"net/http"
)

func HandleStatus(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFiles("templates/status.html"))
			tmpl.Execute(w, nil)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		email := r.FormValue("email")
		if email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		application, err := getApplicationByEmail(db, email)
		if err != nil {
			http.Error(w, "Failed to retrieve application status", http.StatusInternalServerError)
			return
		}

		if application == nil {
			w.Write([]byte("No application found for this email"))
			return
		}

		w.Write([]byte("Application status: " + application.Status))
	}
}

func getApplicationByEmail(db *sql.DB, email string) (*models.Application, error) {
	var application models.Application
	query := `SELECT * FROM applications WHERE email = $1`
	row := db.QueryRow(query, email)
	err := row.Scan(
		&application.ID,
		&application.Name,
		&application.Surname,
		&application.Otchestvo,
		&application.Email,
		&application.Password,
		&application.Number,
		&application.Faculty,
		&application.Snils,
		&application.Passport.Seial,
		&application.Passport.Number,
		&application.Passport.Date,
		&application.Passport.Adress,
		&application.Exams.Russian,
		&application.Exams.Math,
		&application.Exams.Physics,
		&application.Exams.Informatics,
		&application.Status,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &application, nil
}

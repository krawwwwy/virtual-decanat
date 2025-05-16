package handlers

import (
	"applicant-service/models"
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
)

func HandleApply(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl := template.Must(template.ParseFiles("templates/form.html"))
			tmpl.Execute(w, nil)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		var application models.Application
		application.Name = r.FormValue("name")
		application.Surname = r.FormValue("surname")
		application.Otchestvo = r.FormValue("otchestvo")
		application.Email = r.FormValue("email")
		application.Password = r.FormValue("password")
		application.Number = r.FormValue("number")
		application.Faculty = r.FormValue("faculty")
		application.Snils = r.FormValue("snils")

		application.Passport.Seial, _ = strconv.Atoi(r.FormValue("passport_seial"))
		application.Passport.Number, _ = strconv.Atoi(r.FormValue("passport_number"))
		application.Passport.Date = r.FormValue("passport_date")
		application.Passport.Adress = r.FormValue("passport_adress")

		application.Exams.Russian, _ = strconv.Atoi(r.FormValue("russian"))
		application.Exams.Math, _ = strconv.Atoi(r.FormValue("math"))
		application.Exams.Physics, _ = strconv.Atoi(r.FormValue("physics"))
		application.Exams.Informatics, _ = strconv.Atoi(r.FormValue("informatics"))

		err := saveApplicationToDB(db, application)
		if err != nil {
			http.Error(w, "Failed to save application", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Application submitted successfully"))
	}
}

func saveApplicationToDB(db *sql.DB, application models.Application) error {
	query := `INSERT INTO applications (name, surname, otchestvo, email, password, number, faculty, snils, passport_seial, passport_number, passport_date, passport_adress, russian, math, physics, informatics)
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`

	_, err := db.Exec(query,
		application.Name,
		application.Surname,
		application.Otchestvo,
		application.Email,
		application.Password,
		application.Number,
		application.Faculty,
		application.Snils,
		application.Passport.Seial,
		application.Passport.Number,
		application.Passport.Date,
		application.Passport.Adress,
		application.Exams.Russian,
		application.Exams.Math,
		application.Exams.Physics,
		application.Exams.Informatics,
	)

	return err
}

package main

import (
	"applicant-service/handlers"
	"applicant-service/utils"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing the connection: %v", err)
		}
	}(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Println("Error parsing template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/apply", handlers.HandleApply(db))
	http.HandleFunc("/status", handlers.HandleStatus(db))

	fmt.Println("Successfully run container")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

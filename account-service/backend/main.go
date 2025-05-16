package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/account.html"))
		tmpl.Execute(w, nil)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

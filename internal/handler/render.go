package handler

import (
	"html/template"
	"log"
	"net/http"
)

func render(w http.ResponseWriter, data interface{}, filename ...string) {
	tmpl, err := template.ParseFiles(filename...)
	if err != nil {
		log.Print(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Print(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
	}
}

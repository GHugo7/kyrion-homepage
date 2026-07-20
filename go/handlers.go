package main

import (
	"html/template"
	"net/http"

	"github.com/metalblueberry/console"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/html/index.html")
	if err != nil {
		console.Error("Erreur lors du chargement du template index")
		return
	}
	type Services struct {
		Titre       string
		Description string
		Categories  string
	}
	tmpl.Execute(w, Services{Titre: "Hey what's up", Categories: "Moto"})
}

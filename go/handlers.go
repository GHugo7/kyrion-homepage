package main

import (
	"encoding/json"
	"net/http"

	"github.com/metalblueberry/console"
)

func handlerServices(w http.ResponseWriter, r *http.Request) {
	console.Info("Chargement des services")

	type Services struct {
		Titre       string `json:"titre"`
		Description string `json:"description"`
		Categories  string `json:"categories"`
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Services{Titre: "Hey what's up", Categories: "Moto"})
}

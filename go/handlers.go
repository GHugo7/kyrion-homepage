package main

import (
	"encoding/json"
	"net/http"

	"github.com/metalblueberry/console"
)

func filterCat(services []Service) []Category {
	grouped := map[string][]Service{}
	var r []Category

	for _, s := range services {
		grouped[s.Categories] = append(grouped[s.Categories], s)
	}
	for nom, liste := range grouped {
		r = append(r, Category{Nom: nom, Services: liste})
	}
	return r
}

func handlerServices(w http.ResponseWriter, r *http.Request) {
	service, err := getDockerServices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		console.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filterCat(service))
}

package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sort"

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

	sort.Slice(r, func(i, j int) bool {
		return r[i].Nom > r[j].Nom
	})

	return r
}

func handlerServices(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		service, err := getDockerServices(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			console.Error(err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filterCat(service))
	}
}

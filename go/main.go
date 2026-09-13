package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/metalblueberry/console"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	console.Clear()
	console.Info("Starting...")

	containerName, err := getContainerNames()
	if err != nil {
		console.Error(err)
		return
	}

	db, err := ConnectDB(containerName)
	if err != nil {
		console.Error(err)
		return
	}
	defer db.Close()

	router.Get("/api/services", handlerServices(db))
	http.ListenAndServe(":1818", router)
}

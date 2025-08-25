package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"jp-song-trainer/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Маршруты
	r.Post("/api/translate", handlers.Translate)
	r.Get("/api/history", handlers.GetHistory)
	r.Delete("/api/history/{id}", handlers.DeleteHistory)
	r.Put("/api/comments/{id}", handlers.UpdateComment)

	log.Println("😜 Backend running on :8080")
	http.ListenAndServe(":8080", r)
}

package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"jp-song-trainer/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// API-эндпоинты
	r.Route("/api/v1", func(v1 chi.Router) {
		v1.Post("/translate", handlers.Translate)
		v1.Get("/history", handlers.GetHistory)
		v1.Delete("/history/{id}", handlers.DeleteHistory)
		v1.Put("/comments/{id}", handlers.UpdateComment)
	})

	// Статическое подключение Swagger
	r.Get("/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/swagger.yaml")
	})

	// Swagger UI
	// Доступно по адресу: http://localhost:8080/swagger/index.html
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger.yaml"),
	))

	log.Println("😜 Backend running on :8080")
	http.ListenAndServe(":8080", r)
}

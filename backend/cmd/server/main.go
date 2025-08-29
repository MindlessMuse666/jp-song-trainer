package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"jp-song-trainer/internal/config"
	"jp-song-trainer/internal/handler"
	"jp-song-trainer/pkg/database"
	"jp-song-trainer/pkg/logger"
)

func main() {
	// Load config
	cfg := config.Load()

	// Init logger
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Init DB
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// Init router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/health", handler.HealthCheck)

	// Start server
	logger.Info("Starting server", zap.String("port", cfg.Port))
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}

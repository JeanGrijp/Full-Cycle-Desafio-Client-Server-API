package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/JeanGrijp/Full-Cycle-Desafio-Client-Server-API/server/controller"
	"github.com/JeanGrijp/Full-Cycle-Desafio-Client-Server-API/server/internal/database"
)

func main() {
	// Start the server

	db, err := database.NewPostgresConnection()
	if err != nil {
		slog.ErrorContext(context.Background(), "Error connecting to database", "err", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/cotacao", controller.ExchangeRateController(db))
	http.HandleFunc("/exchange", controller.ExchangeController(db))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.ErrorContext(context.Background(), "Error starting the server", "err", err)

	}

	slog.Info("Server started on port 8080")
}

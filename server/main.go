package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/JeanGrijp/Full-Cycle-Desafio-Client-Server-API/server/controller"
)

func main() {
	// Start the server

	http.HandleFunc("/cotacao", controller.ExchangeRateController)
	http.HandleFunc("/exchange", controller.ExchangeController)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.ErrorContext(context.Background(), "Error starting the server", "err", err)

	}

	slog.Info("Server started on port 8080")
}

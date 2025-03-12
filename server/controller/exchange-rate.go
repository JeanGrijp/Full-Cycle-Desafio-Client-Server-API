package controller

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type exchangeRateType struct {
	Usdbrl usdbrl `json:"USDBRL"`
}
type usdbrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type exchangeRateResponse struct {
	Dolar float64 `json:"dolar"`
}

func ExchangeRateController(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		slog.ErrorContext(ctx, "Error creating request", "err", err)
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}
	slog.InfoContext(ctx, "Sending request to awesomeapi")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		slog.ErrorContext(ctx, "Error sending request", "err", err)
		http.Error(w, "Error sending request", http.StatusInternalServerError)
		return
	}
	slog.InfoContext(ctx, "Request sent successfully")
	slog.InfoContext(ctx, "Decoding response")

	defer res.Body.Close()

	var exchangeRate exchangeRateType
	err = json.NewDecoder(res.Body).Decode(&exchangeRate)
	if err != nil {
		slog.ErrorContext(ctx, "Error decoding response", "err", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	slog.InfoContext(ctx, "Response decoded successfully")
	slog.InfoContext(ctx, "The response is", "exchangeRate", exchangeRate)

	dolarStr := exchangeRate.Usdbrl.Bid

	dolar, err := strconv.ParseFloat(dolarStr, 64)
	if err != nil {
		slog.ErrorContext(ctx, "Error getting dolar value", "err", err)
		http.Error(w, "Error getting dolar value", http.StatusInternalServerError)
		return
	}
	// Get the dolar value
	slog.InfoContext(ctx, "Dolar value obtained successfully", "dolar", dolar)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exchangeRateResponse{Dolar: dolar})

	slog.InfoContext(ctx, "Response sent successfully")
}

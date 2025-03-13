package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

// func ExchangeController(w http.ResponseWriter, r *http.Request) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
// 	if err != nil {
// 		slog.ErrorContext(ctx, "Error creating request", "err", err)
// 		http.Error(w, "Error creating request", http.StatusInternalServerError)
// 		return
// 	}
// 	slog.InfoContext(ctx, "Sending request to awesomeapi")

// 	client := http.Client{}

// 	res, err := client.Do(req)
// 	if err != nil {
// 		slog.ErrorContext(ctx, "Error sending request", "err", err)
// 		http.Error(w, "Error sending request", http.StatusInternalServerError)
// 		return
// 	}
// 	slog.InfoContext(ctx, "Request sent successfully")
// 	slog.InfoContext(ctx, "Decoding response")

// 	defer res.Body.Close()

// 	var exchangeRate exchangeRateType
// 	err = json.NewDecoder(res.Body).Decode(&exchangeRate)
// 	if err != nil {
// 		slog.ErrorContext(ctx, "Error decoding response", "err", err)
// 		http.Error(w, "Error decoding response", http.StatusInternalServerError)
// 		return
// 	}

// 	slog.InfoContext(ctx, "Response decoded successfully")
// 	slog.InfoContext(ctx, "The response is", "exchangeRate", exchangeRate)

// 	dolarStr := exchangeRate.Usdbrl.Bid

// 	dolar, err := strconv.ParseFloat(dolarStr, 64)
// 	if err != nil {
// 		slog.ErrorContext(ctx, "Error getting dolar value", "err", err)
// 		http.Error(w, "Error getting dolar value", http.StatusInternalServerError)
// 		return
// 	}
// 	// Get the dolar value
// 	slog.InfoContext(ctx, "Dolar value obtained successfully", "dolar", dolar)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(exchangeRate)

// 	slog.InfoContext(ctx, "Response sent successfully")
// }

func ExchangeController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		defer res.Body.Close()
		slog.InfoContext(ctx, "Request sent successfully", "res", res)
		slog.InfoContext(ctx, "Decoding response")

		var exchangeRate exchangeRateType
		err = json.NewDecoder(res.Body).Decode(&exchangeRate)
		if err != nil {
			slog.ErrorContext(ctx, "Error decoding response", "err", err)
			http.Error(w, "Error decoding response", http.StatusInternalServerError)
			return
		}

		slog.InfoContext(ctx, "Response decoded successfully")

		dolarStr := exchangeRate.Usdbrl.Bid
		dolar, err := strconv.ParseFloat(dolarStr, 64)
		if err != nil {
			slog.ErrorContext(ctx, "Error getting dolar value", "err", err)
			http.Error(w, "Error getting dolar value", http.StatusInternalServerError)
			return
		}

		dbCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		// registrar no banco de dados a cotação atual do dólar na data atual
		_, err = db.ExecContext(dbCtx, `INSERT INTO exchange_rates (bid, created_at) VALUES ($1, $2)`, dolar, time.Now())
		if err != nil {
			slog.ErrorContext(ctx, "Error inserting exchange rate into database", "err", err)
			http.Error(w, "Error inserting exchange rate into database", http.StatusInternalServerError)
			return
		}

		slog.InfoContext(ctx, "Dolar value obtained successfully", "dolar", dolar)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(exchangeRate)
	}
}

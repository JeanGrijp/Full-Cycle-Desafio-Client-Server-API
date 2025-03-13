package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

type exchangeRateType struct {
	Dolar float64 `json:"dolar"`
}

func main() {
	// Start the client

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		slog.ErrorContext(ctx, "Error creating request", "err", err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		slog.ErrorContext(ctx, "Error sending request", "err", err)
		return
	}
	defer res.Body.Close()

	slog.InfoContext(ctx, "Request sent successfully")

	// Decode the response and save the file context inside a text file called cotacao.txt
	// The file should contain the dolar value in the format "Dolar: {}"

	var exchangeRate exchangeRateType
	err = json.NewDecoder(res.Body).Decode(&exchangeRate)
	if err != nil {
		slog.ErrorContext(ctx, "Error decoding response", "err", err)
		return
	}

	slog.InfoContext(ctx, "API response", "raw", exchangeRate)

	if exchangeRate.Dolar == 0 {
		slog.ErrorContext(ctx, "Empty Bid field in response")
		return
	}

	dolar := exchangeRate.Dolar

	slog.InfoContext(ctx, "Dolar value obtained successfully", "dolar", dolar)

	file, err := os.Create("cotacao.txt")
	if err != nil {
		slog.ErrorContext(ctx, "Error creating file", "err", err)
		return
	}

	defer file.Close()

	file.WriteString("Dolar: " + strconv.FormatFloat(dolar, 'f', 4, 64))

	slog.InfoContext(ctx, "File created successfully")
	slog.InfoContext(ctx, "Dolar value saved successfully")

}

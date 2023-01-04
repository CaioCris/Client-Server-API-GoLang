package main

import (
	"Client-Server-API-GoLang/domain"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func getUSDBRLExchangeRate() *domain.ExchangeRate {
	exchangeUrl := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", exchangeUrl, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var data domain.ExchangeRate
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return &data
}

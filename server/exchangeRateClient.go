package main

import (
	"Client-Server-API-GoLang/domain"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func getUSDBRLExchangeRate() (*domain.ExchangeRate, error) {
	exchangeUrl := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", exchangeUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data domain.ExchangeRate
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

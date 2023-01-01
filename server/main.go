package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// O server.go deverá consumir a API contendo o câmbio de Dólar e Real
// no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL
// e em seguida deverá retornar no formato JSON o resultado para o cliente.

// Usando o package "context", o server.go deverá registrar no banco de dados
// SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API
// de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir
// os dados no banco deverá ser de 10ms.

// O endpoint necessário gerado pelo server.go para este desafio será:
// /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

type ExchangeRate struct {
	Usdbrl struct {
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
	} `json:"USDBRL"`
}

func main() {
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
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
	var data ExchangeRate
	err = json.Unmarshal(body, &data)
	println(data.Usdbrl.Code)
}

package main

import (
	"Client-Server-API-GoLang/domain"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.

// O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON).
// Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber
// o resultado do server.go.

// O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}

func main() {
	writeDollarExchangeRateFile()
}

func writeDollarExchangeRateFile() {
	var exchangeRate = getDollarExchangeRate()
	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	_, err = file.WriteString(fmt.Sprintf("Dólar: {%v}", exchangeRate.USDBRL.Bid))

}

func getDollarExchangeRate() *domain.ExchangeRate {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
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
			panic(err)
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

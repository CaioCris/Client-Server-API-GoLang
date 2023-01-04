package server

import (
	"encoding/json"
	"net/http"
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

func ExchangeHandler(write http.ResponseWriter, _ *http.Request) {
	var dollarRate = GetDollarExchangeRate()
	write.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(write).Encode(dollarRate)
	if err != nil {
		panic(err)
	}
}

func ExchangeServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", ExchangeHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

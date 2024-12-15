package api

import (
	"encoding/json"
	"net/http"
)

//name of our API CoinBalance

// parameters that our api endpoint will take
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	//Status 200
	Code int
	//Account balance
	Balance int64
}

// Error response
type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error occured", http.StatusInternalServerError)
	}
)

// go mod init myapimodule
// go mod tidy
// go run cmd/api/main.go

package main

import (
	"encoding/json"
	"net/http"
)

// Authorization type manage response output to all transactions
type Authorization struct {
	Approved      bool     `json:"approved"`
	NewLimit      float64  `json:"newLimit"`
	DeniedReasons []string `json:"deniedReasons"`
}

// TransactionAuth responsible to send json with payment informations
func TransactionAuth(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	json.NewDecoder(r.Body).Decode(&payment)
	// Rules

	var auth Authorization
	json.NewEncoder(w).Encode(auth)
}

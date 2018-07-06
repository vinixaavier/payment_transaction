package main

// Transaction type to manage transaction information
type Transaction struct {
	Merchant        string        `json:"merchant"`
	Amount          string        `json:"amount"`
	Time            string        `json:"time"`
	LastTransaction []Transaction `json:"lastTransaction"`
}

package main

// Payment type to manage received payment information
type Payment struct {
	Account         Account       `json:"account"`
	Transaction     Transaction   `json:"transaction"`
	LastTransaction []Transaction `json:"lastTransaction"`
}

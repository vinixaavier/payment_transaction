package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// PostPayment responsible to send json with payment informations
func PostPayment(w http.ResponseWriter, r *http.Request) {
	var payment Payment
	json.NewDecoder(r.Body).Decode(&payment)
	AboveLimit(payment.Account.Limit, payment.Transaction.Amount)
}

func main() {
	routes := mux.NewRouter()
	routes.HandleFunc("/payment", PostPayment).Methods("POST")

	addr := "localhost:3000"

	if err := http.ListenAndServe(addr, routes); err != nil {
		log.Fatal(err)
	}
}

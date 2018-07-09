package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	routes := mux.NewRouter()
	routes.HandleFunc("/payment", TransactionAuth).Methods("POST")

	addr := "localhost:3000"

	if err := http.ListenAndServe(addr, routes); err != nil {
		log.Fatal(err)
	}

}

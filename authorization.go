package main

import (
	"encoding/json"
	"net/http"
)

// Authorization type manage response output to all transactions
type Authorization struct {
	Approved      bool     `json:"approved"`
	NewLimit      float64  `json:"newLimit,omitempty"`
	DeniedReasons []string `json:"deniedReasons,omitempty"`
}

// NewLimit calculates a new limit of account based in amount trasanction
func NewLimit(limit, amount float64) float64 {
	return limit - amount
}

// TransactionAuth responsible to send json with payment informations
func TransactionAuth(w http.ResponseWriter, r *http.Request) {
	var (
		payment Payment
		auth    Authorization
		reasons []string
	)
	json.NewDecoder(r.Body).Decode(&payment)

	if reason, check := CheckCard(payment.Account.CardIsActive); check != true {
		reasons = append(reasons, reason)
	}

	if reason, check := CheckLimit(payment.Account.Limit, payment.Transaction.Amount); check != true {
		reasons = append(reasons, reason)
	}

	if reason, check := CheckBlacklist(payment.Account.Blacklist, payment.Transaction.Merchant); check != true {
		reasons = append(reasons, reason)
	}

	if reason, check := CheckInterval(3, 2, payment.Transaction.Time, payment.LastTransaction); check != true {
		reasons = append(reasons, reason)
	}

	if reason, check := CheckSameMerchant(10, payment.LastTransaction, payment.Transaction.Merchant); check != true {
		reasons = append(reasons, reason)
	}

	if IsWhitelisted(payment.Account.IsWhitelisted) {
		if reason, check := CheckLimitPercent(90, payment.LastTransaction, payment.Account.Limit, payment.Transaction.Amount); check != true {
			reasons = append(reasons, reason)
		}
	}

	if len(reasons) == 0 {
		auth.Approved = true
		auth.NewLimit = NewLimit(payment.Account.Limit, payment.Transaction.Amount)
	} else {
		auth.Approved = false
		auth.DeniedReasons = reasons
	}

	json.NewEncoder(w).Encode(auth)
}

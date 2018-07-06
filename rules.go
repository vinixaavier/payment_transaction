package main

import "fmt"

// AboveLimit The transaction amount should not be above limit
func AboveLimit(limit string, amount string) {
	if amount > limit {
		fmt.Printf("Not authorized!")
	}
} // return true

// CardActive No transaction should be approved when the card is blocked
func CardActive() {} // return true

// NinetyPercent The first transaction shouldn't be above 90% of the limit
func NinetyPercent() {} // return true

// MoreTenMerchant There should not be more than 10 transactions on the same merchant
func MoreTenMerchant() {} // return true

// MerchantBlacklist Check if merchant of the transaction be in blacklist
func MerchantBlacklist() {} // return true

// TwoMinInterval There should not be more than 3 transactions on a 2 minutes interval
func TwoMinInterval() {}

package main

import (
	"fmt"
	"time"
)

// CheckLimit The transaction amount should not be above limit
func CheckLimit(limit, amount float64) (reason string, approve bool) {
	if limit >= amount {
		return "Approved!", true
	}
	return "The transaction amount above limit", false
} // if true, approved!

// CheckCard No transaction should be approved when the card is blocked
func CheckCard(cardIsActive bool) (reason string, approve bool) {
	if cardIsActive {
		return "Approved!", true
	}
	return "Card is blocked!", false
} // if true, approved!

// CheckLimitPercent The first transaction shouldn't be above 90% of the limit
func CheckLimitPercent(percent float64, lastTransaction []Transaction, limit float64, amount float64) (reason string, approve bool) {
	if len(lastTransaction) == 0 {
		ninetyPercent := limit / 100 * percent
		if amount <= ninetyPercent {
			return "Approved", true
		}
		return "The amount above 90% of the limit!", false
	}
	return "Approved", true
} // if true, approved!

// CheckSameMerchant There should not be more than 10 transactions on the same merchant
func CheckSameMerchant(qtd int, lastTransaction []Transaction, merchant string) (reason string, check bool) {
	if len(lastTransaction) >= qtd {
		var count int
		for _, v := range lastTransaction {
			if merchant == v.Merchant {
				count++
				if count >= qtd {
					return "More than 10 transactions on the same merchant!", false
				}
			}
		}
	}
	return "Approved", true
} // if true, approved

// CheckBlacklist Check if merchant of the transaction is blacklisted
func CheckBlacklist(blacklist []string, merchant string) (reason string, check bool) {
	if len(blacklist) != 0 {
		for _, blacklister := range blacklist {
			if merchant != blacklister {
				continue
			}
			return "This merchant is blacklisted", false
		}
	}
	return "Approved!", true
} // if true, approved

// CheckInterval There should not be more than 3 transactions on a 2 minutes interval
func CheckInterval(qtd int, interval float64, moment string, lastTransaction []Transaction) (reason string, check bool) {
	if len(lastTransaction) >= qtd {
		thirdLastTransaction := lastTransaction[len(lastTransaction)-qtd].Time
		t, _ := time.Parse(time.RFC3339, thirdLastTransaction)
		duration := time.Since(t)
		if duration.Minutes() >= interval {
			return "Approved", true
		}
		return fmt.Sprintf("More than %d transactions on %.0f minutes interval!", qtd, interval), false
	}
	return "Approved", true

} // if true, approved

// IsWhitelisted The account of transaction is whitelisted
func IsWhitelisted(isWhitelisted bool) bool {
	if isWhitelisted {
		return true
	}
	return false
}

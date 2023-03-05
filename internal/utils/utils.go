package utils

import "github.com/cryptowatch_challenge/internal/models"

func Reverse(prices []*models.Price) []*models.Price {
	for i := 0; i < len(prices)/2; i++ {
		j := len(prices) - i - 1
		prices[i], prices[j] = prices[j], prices[i]
	}
	return prices
}

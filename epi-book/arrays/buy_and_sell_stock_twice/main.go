package epi

import "math"

// BuyAndSellStockTwice returns the maximum profit that can be made by buying
// and selling a stock twice
func BuyAndSellStockTwice(prices []int) int {
	maxProfit, minPriceSoFar := 0.0, math.Inf(1)
	firstProfits := make([]float64, len(prices)) // max profit made with first buy and sell by day(inclusive)

	// for each day, find the max profit that can be made by buying and selling
	// []int{12,11,13,9,12,8,14,13,15} -> 10
	// A =  {0, 0, 2, 2, 3,3, 6, 6, 7}
	for i, p := range prices {
		minPriceSoFar = math.Min(minPriceSoFar, float64(p))
		maxProfit = math.Max(maxProfit, float64(p)-minPriceSoFar)
		firstProfits[i] = maxProfit
	}

	// Backward pass: find the max profit from second buy OTD
	// B = {7,7,7,7,7,7,2,2,0}
	currMaxPrice := math.Inf(-1)
	for i := len(prices) - 1; i > 0; i-- {
		currMaxPrice = math.Max(currMaxPrice, float64(prices[i]))
		// M = {7,7,7,9,9,10,5,8,6}
		currProfit := currMaxPrice - float64(prices[i]) + firstProfits[i-1]
		maxProfit = math.Max(maxProfit, currProfit)
	}
	return int(maxProfit)
}

// []int{12,11, 13,  9, 12, 8, 14, 13, 15} -> 10
// A =  {0,  0,  2,  2,  3, 3,  6,  6,  7}
// B =  {7,  7,  7,  7,  7, 7,  2,  2, 0}
// M =  {7,  7,  7,  9,  9,10,  5,  8, 6}

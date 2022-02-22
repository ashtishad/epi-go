package epi

import "testing"

func TestBuyAndSellStockTwice(t *testing.T) {
	for _, test := range []struct {
		prices []int
		profit int
	}{
		{[]int{12, 11, 13, 9, 12, 8, 14, 13, 15}, 10},
	} {
		if got := BuyAndSellStockTwice(test.prices); got != test.profit {
			t.Errorf("BuyAndSellStockTwice(%v) = %v, want %v", test.prices, got, test.profit)
		}
	}
}

package epi

import (
	"reflect"
	"testing"
)

func TestBuyAndSellStock(t *testing.T) {
	testCases := []struct {
		prices []int
		result int
	}{
		{[]int{310, 315, 275, 295, 260, 270, 290, 230, 255, 250}, 30},
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, tc := range testCases {
		var want = tc.result
		got := maxProfit(tc.prices)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Source: %d\n Expected: %d\n   Actual: %d\n",
				tc.prices,
				want,
				got)
		}
	}
}

package leetcode

import (
	"reflect"
	"testing"
)

func Test_(t *testing.T) {
	testCases := []struct {
		nums   []int
		result int
	}{
		{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
		{[]int{1, 2, 3, 2, 2}, 4},
	}

	for _, tc := range testCases {
		var want = tc.result
		got := totalFruit(tc.nums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Source: %d\n Expected: %d\n   Actual: %d\n",
				tc.nums,
				want,
				got)
		}
	}
}

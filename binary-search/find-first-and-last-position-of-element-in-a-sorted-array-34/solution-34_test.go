package leetcode

import (
	"reflect"
	"testing"
)

func Test_(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		res    []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
	}

	for _, tc := range testCases {
		var want = tc.res
		got := searchRange(tc.nums, tc.target)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Source: %d\n Expected: %d\n   Actual: %d\n",
				tc.nums,
				want,
				got)
		}
	}
}

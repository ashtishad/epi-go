package leetcode

import (
	"reflect"
	"testing"
)

func Test_(t *testing.T) {
	testCases := []struct {
		nums   []int
		result []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
		{[]int{1, 1, 2}, []int{1}},
	}

	for _, tc := range testCases {
		var want = tc.result
		var got = findDuplicates(tc.nums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Source: %d\n Expected: %d\n   Actual: %d\n",
				tc.nums,
				want,
				got)
		}
	}
}

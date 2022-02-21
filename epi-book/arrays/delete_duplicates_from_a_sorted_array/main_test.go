package epi

import (
	"reflect"
	"testing"
)

func TestDeleteDuplicatesFromSortedArray(t *testing.T) {
	testCases := []struct {
		nums   []int
		result int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{[]int{2, 3, 5, 5, 7, 11, 11, 11, 13}, 6},
	}

	for _, tc := range testCases {
		var want = tc.result
		got := removeDuplicates(tc.nums)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Source: %d\n Expected: %d\n   Actual: %d\n",
				tc.nums,
				want,
				got)
		}
	}
}

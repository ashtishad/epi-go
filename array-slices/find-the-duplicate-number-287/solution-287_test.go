package leetcode

import "testing"

func Test_(t *testing.T) {
	testCases := []struct {
		nums   []int
		result int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
	}

	for _, tc := range testCases {
		var want = tc.result
		got := findDuplicate(tc.nums)

		if got != want {
			t.Errorf("Source: %d\n Expected: %d\n   Actual: %d\n",
				tc.nums,
				want,
				got)
		}
	}
}

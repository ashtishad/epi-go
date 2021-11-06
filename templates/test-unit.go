package main

import "testing"

func Test_(t *testing.T) {
	testCases := []struct {
		n      int
		result int
	}{
		{5, 4},
		{3, 2},
	}

	for _, tc := range testCases {
		var want = tc.result
		got := FUNCTION(tc.n)

		if got != want {
			t.Errorf("Source: %d\n Expected: %d\n   Actual: %d\n",
				tc.n,
				want,
				got)
		}
	}
}

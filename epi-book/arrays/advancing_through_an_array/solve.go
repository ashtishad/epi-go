package epi

import "math"

func canReachEnd(A []int) bool {
	var farthest, i int
	rp := len(A) - 1

	for i <= farthest && i < rp {
		farthest = int(math.Max(float64(farthest), float64(i+A[i])))
		i++
	}
	return farthest >= rp
}

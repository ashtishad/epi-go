package leetcode

import "math"

func findMaxAverage(nums []int, k int) float64 {
	// var n = len(nums)
	var ws int // window sum
	var max = math.Inf(-1)
	var start = 0

	for end, v := range nums {
		ws += v

		if (end - start + 1) == k {
			// calculate average
			max = math.Max(max, float64(ws)/float64(k))
			ws -= nums[start]
			start++
		}
	}
	return max
}

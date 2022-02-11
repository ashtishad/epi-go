// https://leetcode.com/problems/maximum-product-subarray/
// O(n) time, O(n) space

package leetcode

// [-2,0,-1]
func maxProduct(nums []int) int {
	currMin, currMax, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// multiplied by a negative makes big number smaller, small number bigger
		// so we redefine the extremums by swapping them
		if nums[i] < 0 {
			currMax, currMin = currMin, currMax
		}
		// max/min product for the current number is either the current number itself
		// or the max/min by the previous number times the current one
		currMax = max(nums[i], currMax*nums[i])
		currMin = min(nums[i], currMin*nums[i])

		// the newly computed max value is a candidate for our global result
		res = max(res, currMax)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

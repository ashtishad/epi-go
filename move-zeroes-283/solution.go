// https://leetcode.com/problems/move-zeroes/
// Optimal , Time O(N) , Space 0(1)
// input : [0,1,0,3,12] Output: [1,3,12,0,0]

package movezeroes283

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	lastNonZeroFoundAt := 0
	for curr := range nums {
		if nums[curr] != 0 {
			nums[curr], nums[lastNonZeroFoundAt] = nums[lastNonZeroFoundAt], nums[curr] // swap
			lastNonZeroFoundAt++
		}
	}
}

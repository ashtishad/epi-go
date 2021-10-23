package main

func reverseSlice(nums []int) []int {
	var lp, rp = 0, len(nums) - 1

	for lp < rp {
		nums[lp], nums[rp] = nums[rp], nums[lp]
		lp++
		rp--
	}
	return nums
}

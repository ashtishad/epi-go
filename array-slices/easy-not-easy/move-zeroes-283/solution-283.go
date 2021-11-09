package leetcode

func moveZeroes(nums []int) {
	var n = len(nums)
	if n == 1 {
		return
	}
	// [0,1,0,3,12]

	var idx = 0 // new insertable index

	// shift non zero indexes to the left
	for _, num := range nums {
		if num != 0 {
			nums[idx] = num
			idx++
		}
	}
	// fill the rest with zeros
	for idx < n {
		nums[idx] = 0
		idx++
	}
}

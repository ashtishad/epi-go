package main

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func reverse(nums []int, s int, e int) {
	for s < e {
		swap(nums, s, e)
		s++
		e--
	}
}

package leetcode

func findDuplicate2(nums []int) int {
	for _, num := range nums {
		cur := abs(num)
		if nums[cur] < 0 {
			return cur
		}

		nums[cur] = -nums[cur]
	}
	return -1
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

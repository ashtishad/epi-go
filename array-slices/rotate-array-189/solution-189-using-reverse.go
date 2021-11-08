package leetcode

// Input: nums = [1,2,3,4,5,6,7], k = 3, Out = [5,6,7,1,2,3,4]
// 1. [7,6,5,4,3,2,1] => 2. [(5,6,7),(4),3,2,1] => 3. [5,6,7,(1,2,3,4)]

func rotate(nums []int, k int) {
	var n = len(nums)
	if n == 1 {
		return
	}

	k = k % n
	reverse(nums, 0, n-1) // reverse whole array
	reverse(nums, 0, k-1) // reverse first k elements
	reverse(nums, k, n-1) // reverse last n-k elements
}

func reverse(nums []int, start, end int) {
	for start < end {
		swap(nums, start, end)
		start++
		end--
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

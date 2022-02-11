package leetcode

// Given an integer array nums sorted in non-decreasing order,
// return the squares of each number sorted in non-decreasing order.

// Input: nums = [-4,-1,0,3,10]
// [(4),1,0,3,(10)] -> res = [x,x,x,x,100] -> rp--
// [(4),1,0,(3),10] -> res = [x,x,x,16,100] -> lp++
// [4,(1),0,(3),10] -> res = [x,x,9,16,100] -> rp--
// [4,(1),(0),3,10] -> res = [x,1,9,16,100] -> lp++ -> lp = 2 rp = 2
// [4,1,((0)),3,10] -> res = [0,1,9,16,100] -> loop breaks

func sortedSquares(nums []int) []int {
	// for loop: sort elements in place
	var n = len(nums)
	if n == 1 {
		return []int{nums[0] * nums[0]}
	}
	var lp, rp = 0, n - 1
	var res = make([]int, n) // element will be sorted from right to left
	var rIdx = n - 1

	for lp <= rp {
		if abs(nums[lp]) > abs(nums[rp]) {
			res[rIdx] = nums[lp] * nums[lp]
			lp++
		} else {
			res[rIdx] = nums[rp] * nums[rp]
			rp--
		}
		rIdx--
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

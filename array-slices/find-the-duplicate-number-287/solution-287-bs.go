package leetcode

func findDuplicate(nums []int) int {
	var lp, rp = 1, len(nums)
	var dup = -1
	for lp <= rp {
		var mid = lp + (rp-lp)/2
		var c = count(nums, mid)

		if c > mid {
			rp = mid - 1
			dup = mid
		} else {
			lp = mid + 1
		}
	}
	return dup
}

//  If we pick any one of these numbers and count how many numbers are
//  less than or equal to it, the answer will be equal to that number.
//  Only duplicate number will have more than itself.
func count(nums []int, num int) int {
	var total = 0
	for _, n := range nums {
		if n <= num {
			total++
		}
	}
	return total
}

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Time: O(logN)  Space: O(1)

/*
If A[mid] < target, then the range must begins on the right of mid (hence i = mid+1 for the next iteration)
If A[mid] > target, it means the range must begins on the left of mid (j = mid-1)
If A[mid] = target, then the range must begins on the left of or at mid (j= mid)
Since we would move the search range to the same side for case 2 and 3, we might as well merge them as one single case so that less code is needed:
*/

package leetcode

func searchRange(nums []int, target int) []int {
	var res = []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	var lp, rp = 0, len(nums) - 1
	res[0] = findStartingIndex(nums, target, lp, rp)
	if res[0] == -1 {
		return res
	}

	res[1] = findEndingIndex(nums, target, res[0], rp)

	return res
}

// Input: nums = [5,7,7,8,8,8,8,9,10], target = 8 , Output: 3
func findStartingIndex(nums []int, target int, lp int, rp int) (idx int) {
	idx = -1
	for lp <= rp {
		var mid = lp + (rp-lp)/2
		var midVal = nums[mid]
		// case 1: mid val is larger or equal to target
		// case else : mid val is lower than target
		if midVal >= target {
			rp = mid - 1
		} else {
			lp = mid + 1
		}

		// keep updating target, we will end up with the lowest
		if midVal == target {
			idx = mid
		}
	}
	return idx
}

// Input: nums = [5,7,7,8,8,8,8,9,10], target = 8 , Output: 6
func findEndingIndex(nums []int, target int, lp int, rp int) (idx int) {
	idx = -1
	for lp <= rp {
		var mid = lp + (rp-lp)/2
		var midVal = nums[mid]
		// case 1: mid val is lower or equal to target
		// case else : mid val is higher than target
		if midVal <= target {
			lp = mid + 1
		} else {
			rp = mid - 1
		}

		// keep updating target, we will end up with the highest
		if midVal == target {
			idx = mid
		}
	}
	return idx
}

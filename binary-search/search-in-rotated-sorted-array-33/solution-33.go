// https://leetcode.com/problems/search-in-rotated-sorted-array
// Time: O(logN)  Space: O(1)
// Input: nums = [4,5,6,7,0,1,2], target = 0

package leetcode

func findSmallestElement(nums []int, lp int, rp int) int {
	for lp < rp {
		mid := lp + (rp-lp)/2

		if nums[mid] > nums[rp] {
			lp = mid + 1 // that's how we reach a sorted array
		} else {
			rp = mid // that's normal, then we can use left sub-list
		}
	}
	return lp
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	lp := 0
	rp := len(nums) - 1

	pivot := findSmallestElement(nums, lp, rp) // smallest element

	// defining searching boundary
	if target >= nums[pivot] && target <= nums[rp] {
		lp = pivot
	} else {
		rp = pivot
	}

	// regular binary search
	for lp <= rp {
		mid := lp + (rp-lp)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// use left sub-list
			rp = mid - 1
		} else {
			// use right sub-list
			lp = mid + 1
		}
	}
	return -1
}

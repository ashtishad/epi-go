//  https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/discuss/158940/Beat-100%3A-Very-Simple-(Python)-Very-Detailed-Explanation

package leetcode

func findMin(nums []int) int {
	var n = len(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	// Input: nums = [3,4,5,1,2] Output: 1
	// Explanation: The original array was [1,2,3,4,5] rotated 3 times.
	var lp, rp = 0, n - 1

	// case : No rotation or already in initial state
	if nums[rp] > nums[lp] {
		return nums[lp]
	}

	// case if : that's not normal, so pivot index must be somewhere in right sublist.
	// case else : nums[mid] < nums[right], normal situation,
	// all values are unique, so no == check
	for lp < rp {
		var mid = lp + (rp-lp)/2
		if nums[mid] > nums[rp] {
			lp = mid + 1
		} else {
			rp = mid
			// it is possible for the mid index to store a smaller value
		}
	}
	return nums[lp] // coz, look breaks at left < right, left should have minimum value
}

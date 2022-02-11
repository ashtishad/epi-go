// time complexity: O(nlogn)

package leetcode

func twoSumSearch(numbers []int, target int) []int {
	// Input: numbers = [2,7,11,15], target = 9
	// Output: [1,2]
	for i, num := range numbers {
		var idx = search(numbers, target-num, i+1)
		if idx != -1 {
			return []int{i + 1, idx + 1}
		}
	}
	return nil
}

func search(nums []int, target int, lp int) int {
	var rp = len(nums) - 1

	for lp <= rp {
		mid := lp + (rp-lp)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			rp = mid - 1
		} else {
			lp = mid + 1
		}
	}
	return -1
}

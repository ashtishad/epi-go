package leetcode

import "sort"

func threeSumClosest(nums []int, target int) int {
	var n = len(nums)
	var res = nums[0] + nums[1] + nums[2]
	if n == 3 {
		return res
	}
	sort.Ints(nums)

	for i := range nums {
		// SKIP duplecates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// a fixed, looking for b and c(two sum)
		rp, lp := i+1, n-1
		for rp < lp {
			var sum = nums[i] + nums[rp] + nums[lp]
			// if eligible, update closer sum
			if abs(target-sum) < abs(target-res) {
				res = sum
			}

			if sum <= target {
				rp++
			} else {
				lp--
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

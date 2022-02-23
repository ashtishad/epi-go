package epi

// [2,(3),6,5,4,1] -> 2,(4),6,5,(3),1 -> 2,4, 1,3,5,6
func nextPermutation(nums []int) {
	pidx := checkPermutationPossibility(nums)
	// If the array is already sorted, no need to do anything
	if pidx == -1 {
		reverse(nums, 0, len(nums)-1)
		return
	}

	rp := len(nums) - 1
	// start from right most to leftward,find the first number which is larger than PIVOT
	for rp > 0 {
		if nums[rp] > nums[pidx] {
			nums[pidx], nums[rp] = nums[rp], nums[pidx]
			break
		}
		rp--
	}
	// Finally, Reverse all elements which are right from pivot
	reverse(nums, pidx+1, len(nums)-1)
}

// checkPermutationPossibility returns 1st occurrence Index where
// value is not in increasing order(from right to left)
// returns -1 if not found(it's already in its last permutation)
func checkPermutationPossibility(nums []int) (idx int) {
	// search right to left for 1st number(from right) that is not in increasing order
	var rp = len(nums) - 1
	for rp > 0 {
		if nums[rp-1] < nums[rp] {
			idx = rp - 1
			return idx
		}
		rp--
	}
	return -1
}

func reverse(nums []int, s int, e int) {
	for s < e {
		nums[s], nums[e] = nums[e], nums[s]
		s++
		e--
	}
}

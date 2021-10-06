// Space Optimal, Operation Sub-Optimal
// Time O(N) , Space 0(1)

package movezeroes283

func moveZeroesEasy(nums []int) {
	lastNonZeroFoundAt := 0

	for curr := range nums {
		if nums[curr] != 0 {
			nums[lastNonZeroFoundAt] = nums[curr]
			lastNonZeroFoundAt++
		}
	}

	l := len(nums)
	for i := lastNonZeroFoundAt; i < l; i++ {
		nums[i] = 0
	}
}

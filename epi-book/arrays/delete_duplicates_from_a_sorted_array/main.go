package epi

// O(n) time. O(1) space.
// 26. Remove Duplicates from Sorted Array - Easy
func removeDuplicates(nums []int) int {
	// write index, where unique value will be written
	// after traversing the array, as array index is 0 based
	// write index will be the number of unique elements
	idx := 1

	for _, num := range nums {
		if nums[idx-1] != num {
			nums[idx] = num
			idx++
		}
	}
	return idx
}

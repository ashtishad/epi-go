// https://leetcode.com/problems/contains-duplicate/

package leetcode

func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	seen := map[int]bool{}
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = true
	}
	return false
}

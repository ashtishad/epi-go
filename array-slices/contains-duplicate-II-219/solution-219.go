// https://leetcode.com/problems/contains-duplicate-ii/
// Time : O(N) Space : O(1)

package leetcode

func containsNearbyDuplicate(nums []int, k int) bool {
	var ctr = make(map[int]int, len(nums)) // key: num val: index
	for i, val := range nums {
		if idx, ok := ctr[val]; ok && abs(idx-i) <= k {
			return true
		}
		ctr[val] = i
	}

	return false
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

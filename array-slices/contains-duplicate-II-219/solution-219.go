// https://leetcode.com/problems/contains-duplicate-ii/
// Time : O(N) Space : O(1)

package contains_duplicate_II_219

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func containsNearbyDuplicate(nums []int, k int) bool {
	var seen = make(map[int]int, len(nums)) // key: num val: index
	for currIndex, val := range nums {
		if seenNumIndex, ok := seen[val]; ok && abs(seenNumIndex-currIndex) <= k {
			return true
		}
		seen[val] = currIndex
	}

	return false
}

// 96.89% faster
// Time and Space O(N)
// https://leetcode.com/problems/two-sum/

package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	var result []int

	for i := range nums {
		if val, found := seen[target-nums[i]]; found {
			result = []int{val, i}
			break
		}
		seen[nums[i]] = i
	}
	return result
}

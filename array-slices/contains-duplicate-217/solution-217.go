// https://leetcode.com/problems/contains-duplicate/

package main

import "fmt"

func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	seen := map[int]bool{}
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
			break
		}
		seen[num] = true
	}
	return false
}

func main() {
	inputs := []int{1, 2, 3}
	fmt.Println(containsDuplicate(inputs))
}

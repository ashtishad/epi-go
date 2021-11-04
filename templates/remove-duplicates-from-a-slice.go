package main

//  Input: nums = [-1-1,0,0,1,2,-1,-4]
func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	var res []int
	for _, num := range nums {
		if _, ok := seen[num]; !ok {
			seen[num] = true
			res = append(res, num)
		}
	}
	return res
}

package main

//  Input: nums = [-1-1,0,0,1,2,-1,-4]
func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	var uniqNums []int
	for _, num := range nums {
		if _, ok := seen[num]; !ok {
			seen[num] = true
			uniqNums = append(uniqNums, num)
		}
	}
	return uniqNums
}

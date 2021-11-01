package main

import "fmt"

// 96.89% faster
// Time and Space O(N)
// https://leetcode.com/problems/two-sum/

func twoSumss(nums []int, target int) []int {
	hMap := map[int]int{} // Key : num   value : index
	for i := range nums {
		complement := target - nums[i]
		_, ok := hMap[complement]
		if ok == true {
			return []int{i, hMap[complement]}
		}
		hMap[nums[i]] = i
	}
	return nil
}

func main() {
	inputs := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(inputs, target)
	fmt.Println(res)
}

/*
1. create a  hashmap, where Key: num and value: index
2. Lookup for compliment(target-num[i]) exist in map or not
3. If exists in map, return result array(i, hMap[compliment])
4. If not exists, add this to map
*/

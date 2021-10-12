package main

import "testing"

func twoSum1(nums []int, target int) []int {
	hMap := map[int]int{}
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

func twoSum2(nums []int, target int) []int {
	hMap := map[int]int{}
	for i, val := range nums {
		complement := target - val
		_, ok := hMap[complement]
		if ok == true {
			return []int{i, hMap[complement]}
		}
		hMap[val] = i
	}
	return nil
}

func BenchmarkTwoSum1(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	b.ResetTimer()

	var whatever []int

	for i := 0; i < b.N; i++ {
		whatever = twoSum1(nums, target)
	}

	_ = whatever
}

func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	b.ResetTimer()

	var whatever []int

	for i := 0; i < b.N; i++ {
		whatever = twoSum2(nums, target)
	}

	_ = whatever
}

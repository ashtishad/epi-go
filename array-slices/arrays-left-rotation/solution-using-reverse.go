package main

import (
	"fmt"
)

func rotateLeft(d int, nums []int) []int {
	length := len(nums)
	d %= length

	// Reverse all numbers
	reverse(nums, 0, length-1)

	// Reverse first nums.length-d numbers
	reverse(nums, 0, length-d-1)

	// Reverse last d numbers
	reverse(nums, length-d, length-1)

	return nums
}

func reverse(nums []int, start int, end int) []int {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	result := rotateLeft(9, s)

	fmt.Println(result)
}

/*
reverse all the numbers in the array.
Reverse the first (size - k) numbers.
Reverse the last k numbers.

*/

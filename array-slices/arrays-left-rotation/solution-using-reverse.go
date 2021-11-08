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

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func reverse(nums []int, s int, e int) {
	for s < e {
		swap(nums, s, e)
		s++
		e--
	}
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

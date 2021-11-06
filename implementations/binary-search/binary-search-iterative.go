package main

import (
	"fmt"
)

func binarySearch(nums []int, target int) int {
	var le, ri = 0, len(nums) - 1

	for le <= ri {
		mid := le + (ri-le)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			// use left sub-list
			ri = mid - 1
		} else {
			// use right sub-list
			le = mid + 1
		}
	}

	// If we get here we tried to look at an invalid sub-list
	// which means the number isn't in our list.
	return -1
}

func main() {
	target := 10
	sortedList := []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	index := binarySearch(sortedList, target)
	if index >= 0 {
		fmt.Println("Target found in index : ", index)
	} else {
		fmt.Println("Not found", index)
	}
}

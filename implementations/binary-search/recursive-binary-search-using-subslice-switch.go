package main

import "fmt"

func binarySearchRecursive(a []int, search int) (result int, searchCount int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // not found
	case a[mid] > search:
		result, searchCount = binarySearchRecursive(a[:mid], search)
	case a[mid] < search:
		result, searchCount = binarySearchRecursive(a[mid+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += mid + 1
		}
	default: // a[mid] == search
		result = mid // found
	}
	searchCount++
	return
}

func main() {
	target := 10
	sortedList := []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	index, count := binarySearchRecursive(sortedList, target)
	if index >= 0 {
		fmt.Println("Target found in index : ", index)
		fmt.Println("Total Iteration : ", count)
	} else {
		fmt.Println("Not found", index)
	}
}

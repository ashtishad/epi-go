package main

import (
	"errors"
	"fmt"
	"os"
)

func reverseArray(nums []int, start int, end int) ([]int, error) {
	if end > len(nums)-1 || start < 0 {
		return nil, errors.New("please enter valid index")
	}
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums, nil
}

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	result, err := reverseArray(s, 1, 4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}

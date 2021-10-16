package main

import "fmt"

func maxOfNums(nums []int) int {
	res := 0
	if len(nums) == 0 || nums == nil {
		//panic(nums)
		return res
	}
	for i, num := range nums {
		if i == 0 || num > res {
			res = num
		}
	}
	return res
}

func main() {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, -8, -7, 4}
	var empty []int
	fmt.Println(maxOfNums(inputs))
	fmt.Println(maxOfNums(empty))
}

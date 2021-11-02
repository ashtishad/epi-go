package main

import "fmt"

/*
Optimised Approach

Since in XOR Operations 0^0 = 0 and 1^1 = 0.
Hence here the same bits(numbers) will cancel out each other
i.e. result to zero and only the unique bit(number) will be left.

*/

func singleNumber(nums []int) int {
	var n = len(nums)
	for i := 1; i < n; i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

func main() {
	var in = []int{2, 3, 4, 5, 2, 3, 4}
	fmt.Println(singleNumber(in))
}

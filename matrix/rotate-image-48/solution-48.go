package main

import "fmt"

func reverse(nums []int) []int {
	var lp, rp = 0, len(nums) - 1

	for lp < rp {
		nums[lp], nums[rp] = nums[rp], nums[lp]
		lp++
		rp--
	}
	return nums
}

func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 1 {
		return
	}

	for i := 0; i < n; i++ {
		// transpose matrix
		for j := i + 1; j < n; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
		// reverse each row of the matrix
		matrix[i] = reverse(matrix[i])
	}
}

func main() {
	var matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}

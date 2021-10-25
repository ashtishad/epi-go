package main

import "fmt"

func transposeSquareMatrix(matrix [][]int) {
	var n = len(matrix)
	for i := 0; i < n; i++ {
		// i =row, j= column, diagonal values are not changing
		// swap i,j with j,i, as in (0,1) with (1,0)
		// j =i+1, to avoid diagonal indexes of a square matrix
		for j := i + 1; j < n; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}

func main() {
	var matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	transposeSquareMatrix(matrix)
	fmt.Println(matrix)
}

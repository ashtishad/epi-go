package main

import (
	"fmt"
	"sort"
)

func main() {
	var matrix = [][]int{
		{2, 3},
		{6, 3},
		{1, 4},
	}
	sort.Slice(matrix, func(i, j int) bool { return matrix[i][0] < matrix[j][0] })
	fmt.Println(matrix)
}

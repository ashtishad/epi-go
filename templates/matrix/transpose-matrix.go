package main

func transpose(matrix [][]int) {
	var row = len(matrix)
	var col = len(matrix[0])
	for i := 0; i < row; i++ {
		// i =row, j= column, diagonal values are not changing
		// swap i,j with j,i, as in (0,1) with (1,0)
		// j =i+1, to avoid diagonal indexes
		for j := i + 1; j < col; j++ {
			matrix[i][j] = matrix[j][i]
		}
	}
}

//func main() {
//	var matrix = [][]int{
//		{1, 2, 3, 4},
//		{5, 6, 7, 8},
//		{9, 10, 11, 12},
//		{13, 14, 15, 16},
//	}
//	transpose(matrix)
//	fmt.Println(matrix)
//}

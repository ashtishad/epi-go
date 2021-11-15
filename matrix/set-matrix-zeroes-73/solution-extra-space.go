// Easy, Time: O(M×N) Space : O(M+N)

package leetcode

func setZeroes2(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	var rowList []int
	var colList []int
	// mark the rows and columns that are to be made zero.
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowList = append(rowList, i)
				colList = append(colList, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for _, v := range colList {
			matrix[i][v] = 0
		}
	}

	for j := 0; j < n; j++ {
		for _, v := range rowList {
			matrix[v][j] = 0
		}
	}
}

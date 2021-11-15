// Space optimal, Time: O(M×N) Space : O(1)

package leetcode

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	isRow := false // if isRow needs to be 0'd out

	// find rows/cols to zero out
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0

				if r > 0 {
					matrix[r][0] = 0
				} else {
					isRow = true
				}
			}
		}
	}

	// zero out rows/cols except first row/col
	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if matrix[0][c] == 0 || matrix[r][0] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	// zero out first col if necessary
	if matrix[0][0] == 0 {
		for r := 1; r < rows; r++ {
			matrix[r][0] = 0
		}
	}

	// zero out first row if necessary
	if isRow {
		for c := 0; c < cols; c++ {
			matrix[0][c] = 0
		}
	}
}

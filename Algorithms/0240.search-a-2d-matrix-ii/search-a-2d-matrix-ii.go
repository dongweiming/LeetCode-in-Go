package problem0240

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	width := len(matrix[0])
	height := len(matrix)

	row := height - 1
	col := 0
	for row >= 0 && col < width {
		if matrix[row][col] > target {
			row -= 1
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}
	return false
}

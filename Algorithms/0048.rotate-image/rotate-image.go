package problem0048

func rotate(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	for i := 0; i < row; i++ {
		for j := i + 1; j < col; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][col-j-1]
			matrix[i][col-j-1] = tmp
		}
	}
}

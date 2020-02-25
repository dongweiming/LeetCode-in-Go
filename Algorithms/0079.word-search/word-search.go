package problem0079

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	if len(board[0]) == 0 {
		return false
	}

	if len(word) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if backtrack(board, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func backtrack(board [][]byte, word string, index int, i int, j int) bool {
	tmp := board[i][j]
	if tmp != word[index] {
		return false
	}

	if len(word)-1 == index {
		return true
	}

	board[i][j] = 0
	index++
	if (i > 0 && backtrack(board, word, index, i-1, j)) ||
		(j > 0 && backtrack(board, word, index, i, j-1)) ||
		(i < len(board)-1 && backtrack(board, word, index, i+1, j)) ||
		(j < len(board[0])-1 && backtrack(board, word, index, i, j+1)) {
		return true
	}
	board[i][j] = tmp
	return false
}

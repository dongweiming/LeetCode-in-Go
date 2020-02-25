package problem0200

func numIslands(grid [][]byte) int {
	count := 0
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				DfsMarking(grid, i, j, n, m)
				count++
			}
		}
	}
	return count
}

func DfsMarking(grid [][]byte, i, j, n, m int) {
	if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	DfsMarking(grid, i+1, j, n, m)
	DfsMarking(grid, i-1, j, n, m)
	DfsMarking(grid, i, j+1, n, m)
	DfsMarking(grid, i, j-1, n, m)
}

package problem0221

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([]int, n+1)
	max, prev, min := 0, 0, 2<<32
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if matrix[i-1][j-1] == '1' {
				if dp[j-1] < prev {
					min = dp[j-1]
				} else {
					min = prev
				}
				if min < dp[j] {
					dp[j] = min
				}
				dp[j]++
				if dp[j] > max {
					max = dp[j]
				}
			} else {
				dp[j] = 0
			}
			prev = tmp
		}
	}
	return max * max
}

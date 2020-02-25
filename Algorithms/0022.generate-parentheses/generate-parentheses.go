package problem0022

func generateParenthesis(n int) []string {
	res := []string{}
	backtrack(&res, n, "", 0, 0)
	return res
}

func backtrack(res *[]string, n int, S string, left int, right int) {
	if len(S) == 2*n {
		*res = append(*res, S)
	}
	if left < n {
		backtrack(res, n, S+"(", left+1, right)
	}
	if right < left {
		backtrack(res, n, S+")", left, right+1)
	}
	return
}

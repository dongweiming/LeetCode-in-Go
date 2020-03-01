package problem0461

func hammingDistance(x int, y int) int {
	cnt, xor := 0, x^y
	for xor != 0 {
		cnt++
		xor = xor & (xor - 1)
	}
	return cnt
}

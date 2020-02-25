package problem0279

import (
	"math"
)

func numSquares(n int) int {
	if isSquare(n) {
		return 1
	}
	for n%4 == 0 {
		n >>= 2
	}
	if n%8 == 7 {
		return 4
	}
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}
	return 3
}

func isSquare(n int) bool {
	s := math.Sqrt(float64(n))
	return s == float64(int(s))
}

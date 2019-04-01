package problem0967

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	N   int
	K   int
	ans []int
}{

	{
		3,
		7,
		[]int{181, 292, 707, 818, 929},
	},

	{
		2,
		1,
		[]int{10, 12, 21, 23, 32, 34, 43, 45, 54, 56, 65, 67, 76, 78, 87, 89, 98},
	},

	// 可以有多个 testcase
}

func Test_numsSameConsecDiff(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, numsSameConsecDiff(tc.N, tc.K), "输入:%v", tc)
	}
}

func Benchmark_numsSameConsecDiff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			numsSameConsecDiff(tc.N, tc.K)
		}
	}
}
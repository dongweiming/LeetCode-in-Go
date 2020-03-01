package problem0337

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func rob(root *TreeNode) int {
	res := rob_(root)
	return max(res[0], res[1])
}

func rob_(root *TreeNode) []int {
	res := make([]int, 2)
	if root == nil {
		return res
	}
	left, right := rob_(root.Left), rob_(root.Right)
	res[0] = max(left[0], left[1]) + max(right[0], right[1])
	res[1] = left[0] + right[0] + root.Val
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package problem0104

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var max int
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		max = left
	} else {
		max = right
	}
	return max + 1
}

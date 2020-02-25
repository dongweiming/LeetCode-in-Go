package problem0098

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func isValidBST(root *TreeNode) bool {
	stack, cur := &[]*TreeNode{}, -1<<63
	for len(*stack) != 0 || root != nil {
		for root != nil {
			*stack = append(*stack, root)
			root = root.Left
		}
		n := len(*stack) - 1
		root = (*stack)[n]
		(*stack) = (*stack)[:n]
		if root.Val <= cur {
			return false
		}
		cur = root.Val
		root = root.Right
	}
	return true
}

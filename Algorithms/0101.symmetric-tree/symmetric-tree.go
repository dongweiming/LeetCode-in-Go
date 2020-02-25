package problem0101

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	stack := &[]*TreeNode{}
	*stack = append(*stack, root, root)
	for len(*stack) > 0 {
		n := len(*stack) - 1
		t1 := (*stack)[n-1]
		t2 := (*stack)[n]
		(*stack) = (*stack)[:n-1]
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		*stack = append(*stack, t1.Left, t2.Right)
		*stack = append(*stack, t1.Right, t2.Left)
	}
	return true
}

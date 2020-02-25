package problem0102

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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levels := &[][]int{}
	helper(root, 0, levels)
	return *levels
}

func helper(node *TreeNode, level int, levels *[][]int) {
	if len(*levels) == level {
		*levels = append(*levels, []int{})
	}
	(*levels)[level] = append((*levels)[level], node.Val)
	if node.Left != nil {
		helper(node.Left, level+1, levels)
	}
	if node.Right != nil {
		helper(node.Right, level+1, levels)
	}
}

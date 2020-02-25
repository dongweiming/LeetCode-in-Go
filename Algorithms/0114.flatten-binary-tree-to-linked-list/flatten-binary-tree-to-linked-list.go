package problem0114

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func flatten(root *TreeNode) {
	stack := &[]*TreeNode{}
	*stack = append(*stack, root)
	var pre *TreeNode = nil
	for len(*stack) != 0 {
		n := len(*stack) - 1
		temp := (*stack)[n]
		*stack = (*stack)[:n]
		if pre != nil {
			pre.Right = temp
			pre.Left = nil
		}
		if temp.Right != nil {
			*stack = append(*stack, temp.Right)
		}
		if temp.Left != nil {
			*stack = append(*stack, temp.Left)
		}
		pre = temp
	}
}

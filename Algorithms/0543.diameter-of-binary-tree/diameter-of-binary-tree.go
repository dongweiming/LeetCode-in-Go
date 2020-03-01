package problem0543

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := depth(node.Left)
		right := depth(node.Right)
		if left+right+1 > ans {
			ans = left + right + 1
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	depth(root)
	return ans - 1
}

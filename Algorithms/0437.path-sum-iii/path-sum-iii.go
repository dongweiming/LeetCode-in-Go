package problem0437

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

type TreeNode = kit.TreeNode

func pathSum(root *TreeNode, sum int) int {
	return helper(root, 0, sum, map[int]int{0: 1})
}

func helper(node *TreeNode, curSum, sum int, m map[int]int) int {
	if node == nil {
		return 0
	}
	curSum += node.Val
	summary := m[curSum-sum]
	m[curSum]++
	summary += helper(node.Left, curSum, sum, m)
	summary += helper(node.Right, curSum, sum, m)
	m[curSum]--
	return summary
}

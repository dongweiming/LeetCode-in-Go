package problem0094

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	color int
	node  *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	WHITE, GRAY := 0, 1
	res := &[]int{}
	stack := &[]Stack{
		Stack{WHITE, root},
	}

	for len(*stack) > 0 {
		n := len(*stack) - 1
		s := (*stack)[n]
		*stack = (*stack)[:n]
		if s.node == nil {
			continue
		}
		if s.color == WHITE {
			*stack = append(*stack, Stack{WHITE, s.node.Right})
			*stack = append(*stack, Stack{GRAY, s.node})
			*stack = append(*stack, Stack{WHITE, s.node.Left})
		} else {
			*res = append(*res, s.node.Val)
		}
	}
	return *res
}

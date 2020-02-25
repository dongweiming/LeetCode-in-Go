package problem0148

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

// ListNode 是题目预定义的数据类型
type ListNode = kit.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left, right := sortList(head), sortList(mid)
	h := &ListNode{}
	res := h
	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next, left = left, left.Next
		} else {
			h.Next, right = right, right.Next
		}
		h = h.Next
	}
	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}
	return res.Next
}

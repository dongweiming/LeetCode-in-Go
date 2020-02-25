package problem0206

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

// ListNode 是链接节点
type ListNode struct {
	Val  int
	Next *ListNode
}

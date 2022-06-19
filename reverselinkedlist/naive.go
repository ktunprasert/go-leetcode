package reverselinkedlist

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	nodes := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node.Val)
	}

	var reversed *ListNode = &ListNode{nodes[len(nodes)-1], nil}
	node := reversed
	for i := len(nodes) - 2; i >= 0; i-- {
		node.Next = &ListNode{nodes[i], nil}
		node = node.Next
	}
	return reversed
}

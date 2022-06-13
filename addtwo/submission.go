package addtwo

func GetVal(l *ListNode) int {
	if l == nil {
		return 0
	}
	return l.Val
}

func GetNext(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var node, prev *ListNode
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		sum := GetVal(l1) + GetVal(l2) + carry
		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}

		switch {
		case node == nil:
			node = &ListNode{sum, nil}
			prev = node
		default:
			new := &ListNode{sum, nil}
			prev.Next = new
			prev = new
		}
		l1, l2 = GetNext(l1), GetNext(l2)
	}

	return node
}

package addtwo

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	vals := []int{}
	for p := l; p != nil; p = p.Next {
		vals = append(vals, p.Val)
	}
	return fmt.Sprintf("%v", vals)
}

func (l *ListNode) GetVal() int {
	if l == nil {
		return 0
	}
	return l.Val
}

func (l *ListNode) GetNext() *ListNode {
	if l == nil {
		return nil
	}
	return l.Next
}

func addTwoNumbersEasy(l1 *ListNode, l2 *ListNode) *ListNode {
	var node, prev *ListNode
	var carry int
	for l1 != nil || l2 != nil || carry != 0 {
		sum := l1.GetVal() + l2.GetVal() + carry
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
		l1, l2 = l1.GetNext(), l2.GetNext()
	}

	return node
}

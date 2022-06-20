package mergedlinkedlist

import (
	. "github.com/ktunprasert/go-leetcode/reverselinkedlist"
)

func GetVal(node *ListNode) int {
	if node != nil {
		return node.Val
	}
	return 101
}

func GetNext(node *ListNode) *ListNode {
	if node != nil && node.Next != nil {
		return node.Next
	}
	return nil
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	var merged, node *ListNode

	switch {
	case GetVal(list1) > GetVal(list2):
		merged = &ListNode{GetVal(list2), nil}
		list2 = GetNext(list2)
	default:
		merged = &ListNode{GetVal(list1), nil}
		list1 = GetNext(list1)
	}
	node = merged

	for list1 != nil || list2 != nil {
		switch {
		case GetVal(list1) > GetVal(list2):
			node.Next = &ListNode{GetVal(list2), nil}
			list2 = GetNext(list2)
			node = GetNext(node)
		default:
			node.Next = &ListNode{GetVal(list1), nil}
			list1 = GetNext(list1)
			node = GetNext(node)
		}
	}
	return merged
}

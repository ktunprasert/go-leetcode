package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	node := head
	for i := 1; i < len(nums); i++ {
		node.Next = &ListNode{nums[i], nil}
		node = node.Next
	}
	return head
}

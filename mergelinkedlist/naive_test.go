package mergedlinkedlist

import (
	"testing"

	. "github.com/ktunprasert/go-leetcode/reverselinkedlist"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"e1", args{BuildLinkedList([]int{1, 2, 4}), BuildLinkedList([]int{1, 3, 4})}, BuildLinkedList([]int{1, 1, 2, 3, 4, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTwoLists(tt.args.list1, tt.args.list2)
			for left, right := got, tt.want; left != nil && right != nil; left, right = left.Next, right.Next {
				if left.Val != right.Val {
					t.Errorf("reverseList() = %v, want %v", got.Val, tt.want.Val)
				}
			}
		})
	}
}

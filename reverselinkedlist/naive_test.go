package reverselinkedlist

import (
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"e1", args{buildLinkedList([]int{1, 2, 3, 4, 5})}, buildLinkedList([]int{5, 4, 3, 2, 1})},
		{"e2", args{buildLinkedList([]int{1, 2})}, buildLinkedList([]int{2, 1})},
		{"gotcha", args{buildLinkedList([]int{})}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseList(tt.args.head)
			for left, right := got, tt.want; left != nil && right != nil; left, right = left.Next, right.Next {
				if left.Val != right.Val {

					t.Errorf("reverseList() = %v, want %v", got.Val, tt.want.Val)
				}
			}
			// if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("reverseList() = %v, want %v", got, tt.want)
			// }
		})
	}
}

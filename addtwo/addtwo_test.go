package addtwo

import (
	"fmt"
	"testing"
)

func ListNodeFromArray(arr []int) *ListNode {
	var node, curr *ListNode
	for _, n := range arr {
		if node == nil {
			node = &ListNode{n, nil}
			curr = node
		} else {
			new := &ListNode{n, nil}
			curr.Next = new
			curr = new
		}
	}
	return node
}

func TestAddTwo(t *testing.T) {
	cases := []struct {
		name     string
		inputL1  []int
		inputL2  []int
		expected string
	}{
		{"same_size_arr", []int{2, 4, 3}, []int{5, 6, 4}, "[7 0 8]"},
		{"zero", []int{0}, []int{0}, "[0]"},
		{"uneven", []int{0, 1}, []int{3}, "[3 1]"},
		{"example_3", []int{9, 9, 9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, "[8 9 9 9 0 0 0 0 0 1]"},
	}

	funcs := []struct {
		name string
		fn   func(l1 *ListNode, l2 *ListNode) *ListNode
	}{
		{"easy", addTwoNumbersEasy},
		{"submission", addTwoNumbers},
	}

	for _, f := range funcs {
		for _, tc := range cases {
			t.Run(fmt.Sprintf("%s/%s", f.name, tc.name), func(t *testing.T) {
				l1 := ListNodeFromArray(tc.inputL1)
				l2 := ListNodeFromArray(tc.inputL2)
				result := f.fn(l1, l2)
				if actual := fmt.Sprintf("%v", result); tc.expected != actual {
					t.Fatalf(fmt.Sprintf("Expected %s but got %s", tc.expected, actual))
				}
			})
		}
	}
}

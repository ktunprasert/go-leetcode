package addtwo

import (
	"fmt"
	"testing"
)

func ListNodeFromArray(arr []int) ListNode {
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
	return *node
}

func TestAddTwo_SameSizeArr(t *testing.T) {
	l1 := ListNodeFromArray([]int{2, 4, 3})
	l2 := ListNodeFromArray([]int{5, 6, 4})
	sum := addTwoNumbers(&l1, &l2)
	expected := "[7 0 8]"
	if actual := fmt.Sprintf("%v", sum); expected != actual {
		t.Fatalf(fmt.Sprintf("Expected %s but got %s", expected, actual))
	}
}

func TestAddTwo_Zero(t *testing.T) {
	l1 := ListNodeFromArray([]int{0})
	l2 := ListNodeFromArray([]int{0})
	sum := addTwoNumbers(&l1, &l2)
	expected := "[0]"
	if actual := fmt.Sprintf("%v", sum); expected != actual {
		t.Fatalf(fmt.Sprintf("Expected %s but got %s", expected, actual))
	}
}

func TestAddTwo_UnevenSizeArray(t *testing.T) {
	l1 := ListNodeFromArray([]int{0, 1})
	l2 := ListNodeFromArray([]int{1})
	sum := addTwoNumbers(&l1, &l2)
	expected := "[1 1]"
	if actual := fmt.Sprintf("%v", sum); expected != actual {
		t.Fatalf(fmt.Sprintf("Expected %s but got %s", expected, actual))
	}
}

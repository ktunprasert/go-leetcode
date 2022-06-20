package invertbinarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) String() string {
	out := fmt.Sprintf("%d", t.Val)
	if t.Left != nil {
		out += fmt.Sprintf("%v", t.Left)
	}
	if t.Right != nil {
		out += fmt.Sprintf("%v", t.Right)
	}
	return out
}

func (t *TreeNode) Insert(val int, imba int) *TreeNode {
	switch {
	case t == nil:
		return &TreeNode{val, nil, nil}
	case val*imba > t.Val*imba:
		t.Right = t.Right.Insert(val, imba)
	case val*imba <= t.Val*imba:
		t.Left = t.Left.Insert(val, imba)
	}
	return t
}

func BuildBinaryTree(vals []int, imba int) *TreeNode {
	var tree *TreeNode = &TreeNode{vals[0], nil, nil}
	for i := 1; i < len(vals); i++ {
		tree.Insert(vals[i], imba)
	}
	return tree
}

type args struct {
	root *TreeNode
}

var tests = []struct {
	name string
	args args
	want *TreeNode
}{
	{"e1", args{BuildBinaryTree([]int{4, 2, 7, 1, 3, 6, 9}, 1)}, BuildBinaryTree([]int{4, 7, 2, 9, 6, 3, 1}, -1)},
	{"e2", args{BuildBinaryTree([]int{2, 1, 3}, 1)}, BuildBinaryTree([]int{2, 3, 1}, -1)},
}

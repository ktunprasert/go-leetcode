package invertbinarytree

func Invert(t *TreeNode) *TreeNode {
	switch {
	case t == nil:
		return nil
	default:
		t.Left, t.Right = Invert(t.Right), Invert(t.Left)
	}
	return t
}

func invertTree(root *TreeNode) *TreeNode {
	return Invert(root)
}

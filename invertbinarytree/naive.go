package invertbinarytree

func (t *TreeNode) Invert() *TreeNode {
	switch {
	case t == nil:
		return nil
	default:
		t.Left, t.Right = t.Right.Invert(), t.Left.Invert()
	}
	return t
}

func invertTree(root *TreeNode) *TreeNode {
	return root.Invert()
}

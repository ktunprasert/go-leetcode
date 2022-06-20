package invertbinarytree

func (t *TreeNode) InvertNaive() *TreeNode {
	switch {
	case t == nil:
		return nil
	default:
		t.Left, t.Right = t.Right.InvertNaive(), t.Left.InvertNaive()
	}
	return t
}

func invertTreeNaive(root *TreeNode) *TreeNode {
	return root.InvertNaive()
}

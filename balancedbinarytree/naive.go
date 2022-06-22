package balancedbinarytree

import (
	. "github.com/ktunprasert/go-leetcode/invertbinarytree"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Height(t *TreeNode) int {
	if t == nil {
		return 0
	}
	return 1 + max(Height(t.Left), Height(t.Right))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	height := Height(root.Left) - Height(root.Right)
	if height > 1 || height < -1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

package diameterbintree

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

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)), Height(root.Left)+Height(root.Right))
}

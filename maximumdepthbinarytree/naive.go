package maximumdepthbinarytree

import (
	. "github.com/ktunprasert/go-leetcode/invertbinarytree"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func DepthSearch(t *TreeNode, n int) int {
	if t == nil {
		return n - 1
	}
	return max(DepthSearch(t.Left, n+1), DepthSearch(t.Right, n+1))
}

func maxDepth(root *TreeNode) int {
	return DepthSearch(root, 1)
}

package subtreeoftree

import (
	. "github.com/ktunprasert/go-leetcode/invertbinarytree"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return (p.Val == q.Val) && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == subRoot
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

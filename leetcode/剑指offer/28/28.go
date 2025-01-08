package _8

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return same(root.Left, root.Right)
}

func same(L *TreeNode, R *TreeNode) bool {
	if L == nil {
		return true
	}
	if L != R {
		return false
	}
	if L == R {
		return same(L.Left, R.Left) && same(L.Right, R.Left)
	}
	return true
}

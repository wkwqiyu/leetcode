package _7

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root != nil {
		exchange(root)
		return root
	} else {
		return root
	}
}
func exchange(root *TreeNode) {
	if root.Right != nil {
		exchange(root.Right)
	}
	if root.Left != nil {
		exchange(root.Left)
	}
	root.Right, root.Left = root.Left, root.Right
	return
}

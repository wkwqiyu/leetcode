package _4

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) (res int) {
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root.Right != nil {
			dfs(root.Right)
		}
		k--
		if k == 0 {
			res = root.Val
			return
		}
		if root.Left != nil {
			dfs(root.Left)
		}
	}
	dfs(root)
	return res
}

package _5_2

import "math"

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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := depth(root.Left) - depth(root.Right)
	if res >= -1 && res <= 1 {
		return isBalanced(root.Right) && isBalanced(root.Left) && true
	}
	return isBalanced(root.Right) && isBalanced(root.Left) && false
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(depth(root.Left)), float64(depth(root.Right))) + 1)
}

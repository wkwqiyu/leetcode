package _4

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) [][]int {
	var array [][]int
	var a []int
	if root == nil {
		return array
	}
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root.Left != nil || root.Right != nil {
			a = append(a, root.Val)
		}
		if root.Left != nil {
			dfs(root.Left, target-root.Val)
		}
		if root.Right != nil {
			dfs(root.Right, target-root.Val)
		}
		if root.Right == nil && root.Left == nil && (target-root.Val) == 0 {
			a = append(a, root.Val)
			array = append(array, append([]int(nil), a...))
			fmt.Println(array)
			fmt.Println("cg")
			return
		}
		a = a[:len(a)-1]
		return
	}
	dfs(root, target)
	return array
}

/*func dfs(root *TreeNode, target int, a []int) {
	if root.Left != nil || root.Right != nil {
		a = append(a, root.Val)
	}
	if root.Left != nil {
		dfs(root.Left, target-root.Val, array, a)
	}
	if root.Right != nil {
		dfs(root.Right, target-root.Val, array, a)
	}
	if root.Right == nil && root.Left == nil && (target-root.Val) == 0 {
		a = append(a, root.Val)
		*array = append(*array, append([]int(nil), a...))
		fmt.Println(*array)
		fmt.Println("cg")
		return
	}
	return
}*/

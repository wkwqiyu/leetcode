package _2

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	var a []int
	queue := list.New()
	if root == nil {
		return a
	}
	queue.PushBack(root)
	for queue.Len() != 0 {
		tmp := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		a = append(a, tmp.Val)
		if tmp.Left != nil {
			queue.PushBack(tmp.Left)
		}
		if tmp.Right != nil {
			queue.PushBack(tmp.Right)
		}
	}
	return a
}

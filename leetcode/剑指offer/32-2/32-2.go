package _2_2

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var a [][]int
	queue := list.New()
	if root == nil {
		return a
	}
	queue.PushBack(root)
	for queue.Len() != 0 {
		if queue.Len() != 0 {
			len := queue.Len()
			var curr []int
			for i := 0; i < len; i++ {
				tmp := queue.Front().Value.(*TreeNode)
				queue.Remove(queue.Front())
				curr = append(curr, tmp.Val)
				if tmp.Left != nil {
					queue.PushBack(tmp.Left)
				}
				if tmp.Right != nil {
					queue.PushBack(tmp.Right)
				}
			}
			a = append(a, curr)
		}
		if queue.Len() != 0 {
			len1 := queue.Len()
			var curr1 []int
			for i := 0; i < len1; i++ {
				tmp := queue.Back().Value.(*TreeNode)
				queue.Remove(queue.Back())
				curr1 = append(curr1, tmp.Val)
				if tmp.Left != nil {
					queue.PushFront(tmp.Right)
				}
				if tmp.Right != nil {
					queue.PushFront(tmp.Left)
				}
			}
			a = append(a, curr1)
		}
	}
	return a
}

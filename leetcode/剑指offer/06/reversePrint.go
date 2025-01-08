package _6

/*
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var a *[]int
	a = new([]int)
	Rt(head, a)
	return *a
}

// Rt 递归函数
func Rt(head *ListNode, a *[]int) {
	if head != nil {
		Rt(head.Next, a)
		*a = append(*a, head.Val)
	}
}

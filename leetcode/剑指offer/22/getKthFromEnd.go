package _2

/**
 * Definition for singly-linked list.*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	var target *ListNode
	target = head
	for i := 1; i < k; i++ {
		head = head.Next
	}
	for head.Next != nil {
		head = head.Next
		target = target.Next
	}
	return target
}

package _8

/**
 * Definition for singly-linked list.*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	head1 := head
	if head.Val == val {
		return head.Next
	}
	for head1.Next.Val != val {
		head1 = head1.Next
	}
	head1.Next = head1.Next.Next
	return head
}

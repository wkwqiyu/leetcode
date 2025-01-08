package main

import "fmt"

/**
 * Definition for singly-linked list.*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	head1 := head
	for head1.Next != nil {
		head1 = head1.Next
	} //找到尾结点
	before(head, head.Next)
	return head1
}

func before(l *ListNode, ln *ListNode) {
	if ln.Next != nil {
		before(l.Next, ln.Next)
		ln.Next = l
		return
	}
	//ln.Next = new(ListNode)
	ln.Next = l
	return
}

/*
	func reverseList(head *ListNode) *ListNode {
		var pre *ListNode
		curr := head
		for curr != nil {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		return prev
	}
*/
func main() {
	head := new(ListNode)
	head1 := head
	for i := 1; i <= 5; i++ {
		head.Val = i
		head.Next = new(ListNode)
		head = head.Next
	}
	for head1.Next != nil {
		fmt.Print(head1.Val, " ")
		head1 = head1.Next
	}
	head = reverseList(head1)
	for head.Next != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
}

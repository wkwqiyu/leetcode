package Minstack

import (
	"log"
)

type MinStack struct {
	head *Node
	min  int
}
type Node struct {
	next   *Node
	number int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	point := new(Node)
	point.number = x
	if this.head == nil {
		this.head = point
	} else {
		point.next = this.head
		this.head = point
	}
}

func (this *MinStack) Pop() {
	if this.head != nil {
		this.head = this.head.next
	}
}

func (this *MinStack) Top() int {
	if this.head == nil {
		log.Panic("当前栈是空的")
	}
	return this.head.number
}

func (this *MinStack) Min() int {
	point := this.head
	i := point.number
	point = point.next
	for ; point != nil; point = point.next {
		if i > point.number {
			i = point.number
		}
	}
	return i
}

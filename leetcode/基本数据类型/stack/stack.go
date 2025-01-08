package stackpackage

import (
	"log"
)

type MinStack struct {
	head *Node
}
type Node struct {
	next   *Node
	number interface{}
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

// Push 向栈中放入一个元素
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

// Pop 向栈中取出一个元素
func (this *MinStack) Pop() interface{} {
	value := this.head.number
	if this.head != nil {
		this.head = this.head.next
	}
	return value
}

// Top 取栈顶元素
func (this *MinStack) Top() interface{} {
	if this.head == nil {
		log.Panic("当前栈是空的")
	}
	return this.head.number
}

func (this *MinStack) Min() int {
	point := this.head
	i := point.number.(int)
	point = point.next
	for ; point != nil; point = point.next {
		if i > point.number.(int) {
			i = point.number.(int)
		}
	}
	return i
}

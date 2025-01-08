package _1

import "fmt"

/*
	func validateStackSequences(pushed []int, popped []int) bool {
		pushedlen := len(pushed)
		var stack []int
		stack = make([]int, pushedlen)
		index := 0  //stack的z栈顶元素
		index2 := 0 //用于与栈顶元素比较popped目标下标
		if len(pushed) != 0 {
			stack[index] = pushed[0]
		} else {
			return true
		}
		for i := 0; i < len(pushed); {
			if index >= 0 && stack[index] == popped[i] {
				index--
				i++
			} else {
				index++
				index2++
				if index2 >= len(pushed) {
					return false
				}
				stack[index] = pushed[index2]
			}
		}
		if index != -1 {
			return false
		} else {
			return true
		}
		return false
	}
*/
func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}
	stack := make([]int, len(pushed)) //创建元素
	index1, index2 := 0, 0
	stack[0] = pushed[index1] //压栈
	index1++
	for i := 0; index1 < len(pushed); {
		for i >= 0 && stack[i] == popped[index2] { //栈顶与出栈数相同
			if i == 0 {
				index2++
				break
			}
			index2++
			fmt.Println(index2)
		}
		stack[i] = pushed[index1] //压栈
		i++
		index1++
	}
	if index2 == len(popped) {
		return true
	}
	return false
}

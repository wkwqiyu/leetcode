package _24

import "fmt"

func calculate(s string) int {
	queue := []byte(s)
	var dfs func() int
	dfs = func() int {
		stack := []int{}
		num := 0
		sign := '+'
		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
			if c >= '0' && c <= '9' {
				num = num*10 + int(c-'0')
			}
			if c == '(' {
				num = dfs()
			}
			if c >= '0' && c <= '9' && c != ' ' || len(queue) == 0 {
				if sign == '+' {
					stack = append(stack, num)
					fmt.Println(num)
				} else if sign == '-' {
					stack = append(stack, -num)
					fmt.Println("-", num)
				}
				sign = int32(c)
				num = 0
			}
			if c == ')' {
				break
			}
		}
		var res int
		for _, val := range stack {
			res += val
		}
		return res
	}
	return dfs()
}

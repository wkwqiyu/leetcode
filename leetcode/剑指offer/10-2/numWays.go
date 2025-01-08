package _0_2

import "fmt"

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	a, b, c := 1, 1, 0
	for i := 0; i <= n; i++ {
		a = b
		b = c
		c = (a + b) % 1000000007
		fmt.Println(c)
	}
	return c
}

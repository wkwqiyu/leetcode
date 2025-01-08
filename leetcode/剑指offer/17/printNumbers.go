package _7

import "math"

func printNumbers(n int) []int {
	var c []int
	c = make([]int, int(math.Pow(10, float64(n)))-1)
	end := int(math.Pow(10, float64(n)))
	for i := 1; i < end; i++ {
		c[i-1] = i
	}
	return c
}

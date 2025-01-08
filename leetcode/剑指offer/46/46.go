package _6

import (
	"strconv"
)

func translateNum(num int) int {
	//动态规划，比较第i-1位和第i位的二位数组合如果两位数可以翻译，那么目前的翻译次数是前i-1个不同的翻译次数加i-2个不同的翻译次数
	s := strconv.Itoa(num)
	l := len(s)
	var a string
	var ai int
	i1, i2, c := 1, 1, 0
	for i := 2; i <= l; i++ {
		a = s[i-2 : i]
		ai, _ = strconv.Atoi(a)
		if ai < 26 && ai >= 10 {
			c = i1 + i2
			i1 = i2
			i2 = c
		} else {
			i1 = i2
		}
	}
	return i2
}

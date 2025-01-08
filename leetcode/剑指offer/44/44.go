package _4

import "strconv"

func findNthDigit(n int) int {
	digit := 1                 //数字的位数
	start := 1                 //该位数的起始数字
	count := 9 * digit * start //该位的所有数字占的总位数
	for n > count {
		n -= count
		digit += 1
		start *= 10
		count = 9 * digit * start
	}
	num := start + (n-1)/digit
	s := strconv.Itoa(num)
	res, _ := strconv.Atoi(string(s[(n-1)%digit]))
	return res
}

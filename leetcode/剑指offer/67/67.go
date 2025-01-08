package _7

import (
	"fmt"
	"math"
	"strings"
)

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	var s, res int64
	s = 1
	for index, value := range str {
		if value >= '0' && value <= '9' {
			res = res*10 + int64(value-'0')
			fmt.Println(res)
		} else if value == '-' && index == 0 {
			s = -1
		} else if value == '+' && index == 0 {
		} else {
			break
		}
	}
	fmt.Println(s, res, math.MaxInt32)
	if res > math.MaxInt32 {
		if s == -1 {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	return int(s * res)
}

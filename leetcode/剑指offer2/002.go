package 剑指offer2

import "strconv"

func addBinary(a string, b string) string {
	a1, _ := strconv.ParseInt(a, 2, 64)
	b1, _ := strconv.ParseInt(b, 2, 64)
	a1 += b1
	return strconv.FormatInt(a1, 2)
}

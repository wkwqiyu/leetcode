package _5

func hammingWeight(num uint32) int {
	result := 0
	for ; num > 0; num /= 2 {
		lsb := num % 2
		if lsb == 1 {
			result++
		}
	}
	return result
}

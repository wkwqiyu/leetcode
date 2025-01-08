package _6

func constructArr(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	b := make([]int, len(a))
	b[0] = 1
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}
	temp := 1
	for i := len(a) - 2; i >= 0; i-- {
		temp *= a[i+1]
		b[i] *= temp
	}
	return b
}

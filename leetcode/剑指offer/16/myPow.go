package _6

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	res := float64(1)
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for n > 0 {
		if n%1 == 1 {
			res = res * x
		}
		x = x * x
		n = n >> 1
	}
	return res
}

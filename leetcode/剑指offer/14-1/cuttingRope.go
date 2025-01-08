package _4_1

func cuttingRope(n int) int {
	c := 0
	var i int
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	for i = n; i >= 3; c++ {
		i = i - 3
	}
	switch i {
	case 1:
		{
			c--
			return (PowY(3, c, 1000000007) * 4) % 1000000007
		}
	case 2:
		{
			return (PowY(3, c, 1000000007) * 2) % 1000000007
		}
	case 0:
		return PowY(3, c, 1000000007)
	}
	return 0
}

func PowY(number int, e int, n int) (result int) {
	result = 1
	var a int
	for e > 0 {
		a = e % 2
		if a == 1 {
			result = (number * result) % n
		}
		number = (number * number) % n
		e = e / 2
	}
	return
}

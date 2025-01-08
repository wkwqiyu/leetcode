package _6_1

func singleNumbers(nums []int) []int {
	var num int
	var res []int
	res = make([]int, 2)
	m := 1
	for _, value := range nums {
		num ^= value
	}
	for num&m == 0 {
		m <<= 1
	}
	for _, value := range nums {
		if value&m == 0 {
			res[0] ^= value
		} else {
			res[1] ^= value
		}
	}
	return res
}

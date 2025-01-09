package main

func main() {

	println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]bool)

	for _, value := range nums {
		m[value] = true
	}
	max := 0
	for value, _ := range m {

		if _, ok := m[value-1]; ok {
			continue
		}
		l := 0
		for m[value] {
			value++
			l++
		}
		if max < l {
			max = l
		}
	}
	return max
}

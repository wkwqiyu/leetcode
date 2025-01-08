package _9

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, value := range nums {
		if m[value] < len(nums) {
			m[value]++
		} else {
			return value
		}
	}
	return 0
}

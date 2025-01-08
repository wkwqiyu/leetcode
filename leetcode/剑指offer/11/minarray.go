package _1

func minArray(numbers []int) int {
	n := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if n > numbers[i] {
			return numbers[i]
		}
	}
	return n
}

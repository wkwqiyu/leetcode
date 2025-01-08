package _3_2

func missingNumber(nums []int) int {
	for index, value := range nums {
		if index != value {
			return index
		}
	}
	return len(nums)
}

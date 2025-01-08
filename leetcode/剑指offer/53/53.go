package _3

func search(nums []int, target int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			res++
		}
	}
	return res
}

package _1

func exchange(nums []int) []int {
	c := 0
	for i := 0; i < len(nums); i++ {
		for y := c; y < len(nums); y++ {
			if nums[y]%2 != 0 {
				a := nums[y]
				nums[y] = nums[i]
				nums[i] = a
				y++
				c = y
				break
			}
		}
	}
	return nums
}

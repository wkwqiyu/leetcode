package _3

//first
/*func findRepeatNumber(nums []int) int {
	var n int
	for i := 0; i < len(nums); i++ {
		n = nums[i]
		for y := i + 1; y < len(nums); y++ {
			if n == nums[y] {
				return n
			}
		}
	}
	return 0
}*/

// second
var hash map[int]int

func findRepeatNumber(nums []int) int {
	hash = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if hash[nums[i]] < 0 {
			return nums[i]
		}
		hash[nums[i]] = -(i + 1)
	}
	return 0
}

package _7

func twoSum(nums []int, target int) []int {
	var i, j int
	j = len(nums) - 1
	for i < j {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] == target {
			return []int{nums[i], nums[j]}
		} else {
			j--
		}
	}
	return []int{nums[i], nums[j]}
}

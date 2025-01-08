package _2

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > 0 {
			nums[i+1] += nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	if nums[len(nums)-1] > max {
		max = nums[len(nums)-1]
	}
	return max
}

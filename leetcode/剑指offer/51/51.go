package _1

func reversePairs(nums []int) int {

	var mergesort func(nums []int, start int, end int) int

	mergesort = func(nums []int, start int, end int) int {
		//1分
		if start >= end {
			return 0
		}
		mid := (start + end) / 2
		res := mergesort(nums, start, mid) + mergesort(nums, mid+1, end)
		//2治
		i, j := start, mid
		tmp := []int{}
		for i <= mid && j <= end {
			if nums[i] < nums[j] {
				tmp = append(tmp, nums[i])
				res +=
			}
		}
	}

}

package _1

import (
	"sort"
)

func isStraight(nums []int) bool {
	sort.Ints(nums)
	var joker int
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++
		} else if nums[i]-nums[i+1] != -1 {
			return false
		}
	}
	return nums[4]-nums[joker] < 5
}

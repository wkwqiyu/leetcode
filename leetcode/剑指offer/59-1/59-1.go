package _9_1

import "sort"

type lh struct {
	sort.IntSlice
}

func (l *lh) Push(x interface{}) {
	l.IntSlice = append(l.IntSlice, x.(int))
}
func (l *lh) Pop() interface{} {
	a := l.IntSlice[len(l.IntSlice)-1]
	l.IntSlice = l.IntSlice[0:len(l.IntSlice)]
	return a
}
func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		for len(queue) != 0 && nums[i] > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		if i > k && queue[0] == nums[i-k] { //检测最大值是不是出去了(不影响相等最大值数，新进相等最大值在queue[1])
			queue = queue[1:]
		}
		if i > k-1 {
			res = append(res, queue[0])
		}
	}
	return res
}

package main

import "fmt"

func main() {
	moveZeroesV2([]int{1, 0, 4, 2, 3, 0, 0, 0, 12})
}
func moveZeroes(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if j == i {
				j++
			} else {
				nums[j] = nums[i]
				j++
			}
		}
	}
	for ; j < len(nums); j++ {
		nums[j] = 0
	}
}

/*
如果不出现0元素，
i和j一定相等，这个很重要
所以出现0之后，i一定是指向左侧第一个0的
这时候找下一个不是0的元素和i（左边第一个0）交换
*/
func moveZeroesV2(nums []int) {
	var i, j int
	for j < len(nums) {
		if nums[j] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			//打印nums
			for _, v := range nums {
				fmt.Print(v, " ")
			}
			fmt.Println()
			i++
		}
		j++
	}
	return
}

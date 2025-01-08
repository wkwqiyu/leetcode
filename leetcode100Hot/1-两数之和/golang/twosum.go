package main

/*
两个for循环解决
*/
func twoSum(nums []int, target int) []int {
	for i, valuei := range nums {
		for j := i + 1; j < len(nums); j++ {
			valuej := nums[j]
			if valuei+valuej == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

/*
没有充分利用第一个for循环的产出
*/

/*
改进版本，使用一个map去存储便利过的数据，当便利到下一个数据时与判断map中是否存在m[target-valuei]，存在就证明有target，并返回下标
*/
func twoSumV2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, valuei := range nums {
		if value, ok := m[target-valuei]; ok {
			return []int{i, value}
		}
		m[valuei] = i
	}
	return nil
}

func main() {
	nums := []int{-3, 4, 3, 90}
	target := 0
	twoSum(nums, target)
}

package _7_2

func findContinuousSequence(target int) [][]int {
	if target < 3 {
		return nil
	}
	nums := make([]int, (target/2)+1)
	var res [][]int
	for i := 0; i < (target/2)+1; i++ {
		nums[i] = i + 1
	}
	//fmt.Println(nums)
	i, j := 0, 1
	s := nums[i] + nums[j]
	for i < j {
		//fmt.Println(i,j,s)
		if s == target {
			//fmt.Println("app")
			res = append(res, nums[i:j+1])
			s -= nums[i]
			i++
		} else if s < target {
			j++
			s += nums[j]
		} else if s > target {
			s -= nums[i]
			i++
		}
	}
	return res
}

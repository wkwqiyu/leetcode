package _7

func maxValue(grid [][]int) int {
	//动态规划
	//将该格的上和右中的最大值加到当前格中可以得到该格路径下的最大值
	//从第一排开始加知道最后一排的格全部加完可得最大值
	for index, value := range grid {
		for index1, _ := range value {
			//第一行的情况
			if index == 0 {
				if index1 == 0 {
					continue
				}
				grid[index][index1] += grid[index][index1-1]
				continue
			}
			if index1 == 0 {
				grid[index][index1] += grid[index-1][index1]
				continue
			}
			if grid[index-1][index1] < grid[index][index1-1] {
				grid[index][index1] += grid[index][index1-1]
			} else {
				grid[index][index1] += grid[index-1][index1]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

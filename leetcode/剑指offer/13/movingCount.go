package _3

import "fmt"

func movingCount(m int, n int, k int) int {
	var number *int
	number = new(int)
	var lane [][]int
	lane = make([][]int, m)
	for i := 0; i < m; i++ {
		lane[i] = make([]int, n)
	}
	dfs(0, 0, m, n, k, lane, number)
	return *number
}

func dfs(x, y int, m, n int, k int, lane [][]int, number *int) {
	if x < 0 || x >= m || y < 0 || y >= n {
		return
	}
	x1 := (x / 100) + ((x / 10) % 10) + (x % 10)
	y1 := (y / 100) + ((y / 10) % 10) + (y % 10)
	if x1+y1 > k {
		return
	}
	if lane[x][y] != 1 {
		fmt.Println(lane[x][y], "+1")
		*number++
		lane[x][y] = 1
		dfs(x+1, y, m, n, k, lane, number)
		/*dfs(x-1, y, m, n, k, lane, number)//仔细思考的话，只需要每次向右和下运动
		dfs(x, y+1, m, n, k, lane, number)*/
		dfs(x, y-1, m, n, k, lane, number)
	} else {
		return
	}
}

package _3

import "fmt"

func maxProfit(prices []int) int {
	min := prices[0]
	var dm int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < min {
			min = prices[i]
		}
		fmt.Println(min, dm, prices[i+1]-min)
		if dm < prices[i+1]-min {
			dm = prices[i+1] - min
		}
	}
	return dm
}

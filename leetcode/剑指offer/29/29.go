package _9

import "fmt"

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows, columns := len(matrix), len(matrix[0])
	var order = make([]int, rows*columns)
	var index = 0
	left, right, top, bottom := 0, columns-1, 0, rows-1
	for left <= right && top <= bottom {
		for colum := left; colum <= right; colum++ {
			order[index] = matrix[top][colum]
			index++
		}
		for colum := top + 1; colum <= bottom; colum++ {
			fmt.Println(rows, bottom, colum)
			order[index] = matrix[colum][right]
			index++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				order[index] = matrix[bottom][column]
				index++
			}
			for row := bottom; row > top; row-- {
				fmt.Println(index)
				order[index] = matrix[row][left]
				index++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return order
}

package _5

import (
	"strconv"
)

func minNumber(nums []int) string {
	quicksort(0, len(nums)-1, &nums)
	var res string
	for _, value := range nums {
		res += strconv.Itoa(value)
	}
	return res
}
func quicksort(l, r int, arr *[]int) {
	if l >= r {
		return
	}
	pivot := strconv.Itoa((*arr)[l])
	L, R := l, r
	xy, _ := strconv.Atoi(pivot + strconv.Itoa((*arr)[r]))
	yx, _ := strconv.Atoi(strconv.Itoa((*arr)[r]) + pivot)
	for l < r {
		for l < r && xy <= yx {
			r--
			xy, _ = strconv.Atoi(pivot + strconv.Itoa((*arr)[r]))
			yx, _ = strconv.Atoi(strconv.Itoa((*arr)[r]) + pivot)
		}
		(*arr)[l] = (*arr)[r]
		xy, _ = strconv.Atoi(pivot + strconv.Itoa((*arr)[l]))
		yx, _ = strconv.Atoi(strconv.Itoa((*arr)[l]) + pivot)
		for l < r && xy <= yx {
			l++
			xy, _ = strconv.Atoi(pivot + strconv.Itoa((*arr)[l]))
			yx, _ = strconv.Atoi(strconv.Itoa((*arr)[l]) + pivot)
		}
		(*arr)[r] = (*arr)[l]
	}
	(*arr)[l], _ = strconv.Atoi(pivot)
	quicksort(L, l-1, arr)
	quicksort(r+1, R, arr)
}

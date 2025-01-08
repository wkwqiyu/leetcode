package _0_quick_sort

func getLeastNumbers(arr []int, k int) []int {
	go quicksort(0, len(arr)-1, &arr)
	return arr[0:k]
}

func quicksort(l, r int, arr *[]int) {
	if l >= r {
		return
	}
	pivot := (*arr)[l]
	L, R := l, r
	for l < r {
		for l < r && (*arr)[r] > pivot {
			r--
		}
		(*arr)[l] = (*arr)[r]
		for l < r && (*arr)[l] <= pivot {
			l++
		}
		(*arr)[r] = (*arr)[l]
	}
	(*arr)[l] = pivot
	quicksort(L, l-1, arr)
	quicksort(r+1, R, arr)
}

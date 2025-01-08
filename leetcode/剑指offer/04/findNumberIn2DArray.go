package _4

func findNumberIn2DArray(matrix [][]int, target int) bool {
	b := false

	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) > 0 {
			if target <= matrix[i][len(matrix[i])-1] {
				if target == matrix[i][len(matrix[i])-1] {
					return true
				}
				b = findNumberIn1DArray(matrix[i], target)
				if b == false {
					continue
				} else {
					break
				}
			}
		}
	}
	return b
}

//遍历
/*func findNumberIn1DArray(matrix []int, target int) bool {
	for _, value := range matrix {
		if target == value {
			return true
		}
	}
	return false
}*/
//二分查找
func findNumberIn1DArray(matrix []int, target int) bool {
	high := len(matrix) - 1
	low := 0
	for low <= high {
		mid := (high + low) / 2
		if matrix[mid] == target {
			return true
		} else if matrix[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

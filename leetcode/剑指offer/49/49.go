package _9

import "math"

func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	d := make([]int, n)
	d[0] = 1
	for i := 1; i < n; i++ {
		n2, n3, n5 := d[a]*2, d[b]*3, d[c]*5                                  //做丑数
		d[i] = int(math.Min(math.Min(float64(n2), float64(n3)), float64(n5))) //比大小，将最小的放入当前数组空位中
		if d[i] == n2 {                                                       //谁中标了对应指针往前移动一位
			a++
		}
		if d[i] == n3 {
			b++
		}
		if d[i] == n5 {
			c++
		}
	}
	return d[n-1]
}

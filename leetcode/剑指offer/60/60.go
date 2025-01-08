package _0

import (
	"fmt"
)

func dicesProbability(n int) []float64 {
	f := []float64{float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6), float64(1) / float64(6)}
	for i := 1; i < n; i++ {
		a := make([]float64, i*6)
		for j := 0; j < len(f); j++ {
			for k := j; k < j+6 && k < 6*i; k++ {
				a[k] += f[j] / float64(6)
			}
		}
		f = a
		fmt.Println(f)
	}
	return f
}

// twosum_test.go
package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{-3, 4, 3, 90}, 0, []int{0, 2}},
		//{[]int{3, 2, 4}, 6, []int{1, 2}},
		//{[]int{3, 3}, 6, []int{0, 1}},
		//{[]int{1, 2, 3, 4, 5}, 10, []int{}},
		//{[]int{10, 20, 30, 40, 50}, 100, []int{}},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

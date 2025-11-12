// Package common: util functions
package common

import "fmt"

// SliceArrayEqual Compares if two dimension arrays has same elements (can be different order), each element ([]int) must be the same order
func SliceArrayEqual(x, y [][]int) bool {
	if len(x) != len(y) {
		return false
	}

	sliceToStr := func(s []int) string {
		return fmt.Sprint(s)
	}

	count := make(map[string]int)
	for _, xm := range x {
		count[sliceToStr(xm)]++
	}

	for _, ym := range y {
		key := sliceToStr(ym)
		count[key]--
		if count[key] < 0 {
			return false
		}
	}

	return true
}

func SliceEqual(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	count := make(map[int]int)
	for _, v := range x {
		count[v]++
	}

	for _, v := range y {
		count[v]--
		if count[v] < 0 {
			return false
		}
	}

	return true
}

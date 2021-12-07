package utils

import "math"

func IsIncluded(data []int, itm int) bool {
	for _, i := range data {
		if i == itm {
			return true
		}
	}
	return false
}

func MinMax(data []int64) (int64, int64) {
	min := int64(math.MaxInt64)
	max := int64(math.MinInt64)
	for _, num := range data {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min, max
}

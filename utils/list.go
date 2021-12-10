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

func IsIncluded64(data []int64, itm int64) bool {
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

func Unique(data []int) []int {
	keys := map[int]bool{}
	uniq := []int{}
	for _, itm := range data {
		if _, ok := keys[itm]; !ok {
			keys[itm] = true
			uniq = append(uniq, itm)
		}
	}
	return uniq
}

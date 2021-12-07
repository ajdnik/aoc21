package utils

import (
	"strconv"
	"strings"
)

func ToInt(data string) (int64, error) {
	i, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return i, err
	}
	return i, nil
}

func NormalizeSpaces(data string) string {
	return strings.Join(strings.Fields(data), " ")
}

func BinaryToInt(data string) (int64, error) {
	i, err := strconv.ParseInt(data, 2, 64)
	if err != nil {
		return i, err
	}
	return i, nil
}

func ToIntList(data, delim string) ([]int64, error) {
	parts := strings.Split(data, delim)
	nums := []int64{}
	for _, part := range parts {
		num, err := ToInt(part)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

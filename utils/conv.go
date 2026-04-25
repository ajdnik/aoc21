package utils

import (
	"strconv"
	"strings"
)

func ToInt(data string) (int, error) {
	return strconv.Atoi(data)
}

func NormalizeSpaces(data string) string {
	return strings.Join(strings.Fields(data), " ")
}

func ToIntList(data, delim string) ([]int, error) {
	parts := strings.Split(data, delim)
	nums := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}
	return nums, nil
}

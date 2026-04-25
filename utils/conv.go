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

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ParseDigitGrid(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, ch := range line {
			grid[i][j] = int(ch - '0')
		}
	}
	return grid
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

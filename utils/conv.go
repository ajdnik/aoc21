package utils

import (
	"strconv"
	"strings"
)

// ToInt converts a string to an integer.
func ToInt(data string) (int, error) {
	return strconv.Atoi(data)
}

// NormalizeSpaces collapses all runs of whitespace into single spaces.
func NormalizeSpaces(data string) string {
	return strings.Join(strings.Fields(data), " ")
}

// Abs returns the absolute value of n.
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// ParseDigitGrid converts lines of single-digit characters into a 2D int grid.
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

// ToIntList splits a string by the given delimiter and converts each part to an integer.
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

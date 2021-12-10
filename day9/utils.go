package day9

import (
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

func ToHeights(data string) ([]int64, error) {
	parts := strings.Split(data, "")
	heights := []int64{}
	for _, part := range parts {
		height, err := utils.ToInt(part)
		if err != nil {
			return nil, err
		}
		heights = append(heights, height)
	}
	return heights, nil
}

func FindNeighbors(data [][]int64, row, col int, predicate func(itm int64) bool) []int64 {
	result := []int64{}
	if row+1 < len(data) && predicate(data[row+1][col]) {
		result = append(result, data[row+1][col])
	}
	if row-1 >= 0 && predicate(data[row-1][col]) {
		result = append(result, data[row-1][col])
	}
	if col+1 < len(data[row]) && predicate(data[row][col+1]) {
		result = append(result, data[row][col+1])
	}
	if col-1 >= 0 && predicate(data[row][col-1]) {
		result = append(result, data[row][col-1])
	}
	return result
}

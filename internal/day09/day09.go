package day09

import (
	"slices"

	"github.com/ajdnik/aoc21/utils"
)

func findNeighbors(data [][]int, row, col int, predicate func(itm int) bool) []int {
	var result []int
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

func Part1(lines []string) int {
	heights := utils.ParseDigitGrid(lines)
	var sum int
	for row := range heights {
		for col := range heights[row] {
			res := findNeighbors(heights, row, col, func(itm int) bool {
				return itm <= heights[row][col]
			})
			if len(res) == 0 {
				sum += heights[row][col] + 1
			}
		}
	}
	return sum
}

const noBasin = -1

func mergeBasins(basins [][]int, indexes []int) ([][]int, int) {
	primary := indexes[0]
	remain := indexes[1:]
	for row := range basins {
		for col := range basins[row] {
			if slices.Contains(remain, basins[row][col]) {
				basins[row][col] = primary
			}
		}
	}
	return basins, primary
}

func Part2(lines []string) int {
	heights := utils.ParseDigitGrid(lines)

	basins := make([][]int, len(heights))
	for row := range heights {
		basins[row] = make([]int, len(heights[row]))
		for col := range heights[row] {
			basins[row][col] = noBasin
		}
	}

	var count int
	for row := range heights {
		for col := range heights[row] {
			if heights[row][col] == 9 {
				continue
			}
			neighbours := findNeighbors(basins, row, col, func(itm int) bool {
				return itm != noBasin
			})
			switch len(neighbours) {
			case 0:
				basins[row][col] = count
				count++
			case 1:
				basins[row][col] = neighbours[0]
			default:
				basins, idx := mergeBasins(basins, neighbours)
				basins[row][col] = idx
			}
		}
	}

	sizes := map[int]int{}
	for row := range heights {
		for col := range heights[row] {
			if basins[row][col] == noBasin {
				continue
			}
			sizes[basins[row][col]]++
		}
	}

	var basinSizes []int
	for _, val := range sizes {
		basinSizes = append(basinSizes, val)
	}
	slices.SortFunc(basinSizes, func(a, b int) int { return b - a })
	return basinSizes[0] * basinSizes[1] * basinSizes[2]
}

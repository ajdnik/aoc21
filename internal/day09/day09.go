package day09

import (
	"slices"
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

func toHeights(data string) []int64 {
	parts := strings.Split(data, "")
	heights := make([]int64, len(parts))
	for i, part := range parts {
		height, err := utils.ToInt(part)
		if err != nil {
			panic(err)
		}
		heights[i] = height
	}
	return heights
}

func parseHeights(lines []string) [][]int64 {
	heights := make([][]int64, len(lines))
	for i, line := range lines {
		heights[i] = toHeights(line)
	}
	return heights
}

func findNeighbors(data [][]int64, row, col int, predicate func(itm int64) bool) []int64 {
	var result []int64
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

func Part1(lines []string) int64 {
	heights := parseHeights(lines)
	var sum int64
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[row]); col++ {
			res := findNeighbors(heights, row, col, func(itm int64) bool {
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

func mergeBasins(basins [][]int64, indexes []int64) ([][]int64, int64) {
	primary := indexes[0]
	remain := indexes[1:]
	for row := 0; row < len(basins); row++ {
		for col := 0; col < len(basins[row]); col++ {
			if slices.Contains(remain, basins[row][col]) {
				basins[row][col] = primary
			}
		}
	}
	return basins, primary
}

func Part2(lines []string) int64 {
	heights := parseHeights(lines)

	basins := make([][]int64, len(heights))
	for row := 0; row < len(heights); row++ {
		basins[row] = make([]int64, len(heights[row]))
		for col := 0; col < len(heights[row]); col++ {
			basins[row][col] = noBasin
		}
	}

	var count int64
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[row]); col++ {
			if heights[row][col] == 9 {
				continue
			}
			neighbours := findNeighbors(basins, row, col, func(itm int64) bool {
				return itm != noBasin
			})
			if len(neighbours) == 0 {
				basins[row][col] = count
				count++
			} else if len(neighbours) == 1 {
				basins[row][col] = neighbours[0]
			} else {
				basins, idx := mergeBasins(basins, neighbours)
				basins[row][col] = idx
			}
		}
	}

	sizes := map[int64]int64{}
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[row]); col++ {
			if basins[row][col] == noBasin {
				continue
			}
			sizes[basins[row][col]]++
		}
	}

	var basinSizes []int64
	for _, val := range sizes {
		basinSizes = append(basinSizes, val)
	}
	slices.SortFunc(basinSizes, func(a, b int64) int { return int(b - a) })
	return utils.Mul(basinSizes[0:3])
}

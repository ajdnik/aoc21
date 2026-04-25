package day09

import (
	"slices"
	"strconv"
	"strings"
)

func toHeights(data string) []int {
	parts := strings.Split(data, "")
	heights := make([]int, len(parts))
	for i, part := range parts {
		height, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		heights[i] = height
	}
	return heights
}

func parseHeights(lines []string) [][]int {
	heights := make([][]int, len(lines))
	for i, line := range lines {
		heights[i] = toHeights(line)
	}
	return heights
}

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
	heights := parseHeights(lines)
	var sum int
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[row]); col++ {
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
	for row := 0; row < len(basins); row++ {
		for col := 0; col < len(basins[row]); col++ {
			if slices.Contains(remain, basins[row][col]) {
				basins[row][col] = primary
			}
		}
	}
	return basins, primary
}

func Part2(lines []string) int {
	heights := parseHeights(lines)

	basins := make([][]int, len(heights))
	for row := 0; row < len(heights); row++ {
		basins[row] = make([]int, len(heights[row]))
		for col := 0; col < len(heights[row]); col++ {
			basins[row][col] = noBasin
		}
	}

	var count int
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[row]); col++ {
			if heights[row][col] == 9 {
				continue
			}
			neighbours := findNeighbors(basins, row, col, func(itm int) bool {
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

	sizes := map[int]int{}
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[row]); col++ {
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

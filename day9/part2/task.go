package main

import (
	"log"
	"sort"

	"github.com/ajdnik/aoc21/day9"
	"github.com/ajdnik/aoc21/utils"
)

const NoBasin = -1

func GetNeighbourBasins(basins [][]int, row, col int) []int {
	neighbours := []int{}
	if row+1 < len(basins) && basins[row+1][col] != NoBasin {
		neighbours = append(neighbours, basins[row+1][col])
	}
	if row-1 >= 0 && basins[row-1][col] != NoBasin {
		neighbours = append(neighbours, basins[row-1][col])
	}
	if col+1 < len(basins[row]) && basins[row][col+1] != NoBasin {
		neighbours = append(neighbours, basins[row][col+1])
	}
	if col-1 >= 0 && basins[row][col-1] != NoBasin {
		neighbours = append(neighbours, basins[row][col-1])
	}
	return utils.Unique(neighbours)
}

func MergeBasins(basins [][]int, indexes []int) ([][]int, int) {
	primary := indexes[0]
	remain := indexes[1:]
	for i := 0; i < len(basins); i++ {
		for j := 0; j < len(basins[i]); j++ {
			if utils.IsIncluded(remain, basins[i][j]) {
				basins[i][j] = primary
			}
		}
	}
	return basins, primary
}

func FindBasinSizes(heights [][]int64) []int64 {
	basinMap := make([][]int, len(heights))
	for i := 0; i < len(heights); i++ {
		basinMap[i] = make([]int, len(heights[i]))
		for j := 0; j < len(heights[i]); j++ {
			basinMap[i][j] = NoBasin
		}
	}
	var basinsCount int
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[row]); col++ {
			if heights[row][col] == 9 {
				continue
			}
			neighbours := GetNeighbourBasins(basinMap, row, col)
			if len(neighbours) == 0 {
				basinMap[row][col] = basinsCount
				basinsCount++
			} else if len(neighbours) == 1 {
				basinMap[row][col] = neighbours[0]
			} else {
				basinMap, basinIdx := MergeBasins(basinMap, neighbours)
				basinMap[row][col] = basinIdx
			}
		}
	}
	sizeCount := map[int]int64{}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if basinMap[i][j] == NoBasin {
				continue
			}
			if _, ok := sizeCount[basinMap[i][j]]; !ok {
				sizeCount[basinMap[i][j]] = 0
			}
			sizeCount[basinMap[i][j]]++
		}
	}
	basinSizes := []int64{}
	for _, val := range sizeCount {
		basinSizes = append(basinSizes, val)
	}
	return basinSizes
}

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	heightMap := [][]int64{}
	for scanner.Scan() {
		data := scanner.Text()
		heights, err := day9.ToHeights(data)
		if err != nil {
			log.Fatal(err)
		}
		heightMap = append(heightMap, heights)
	}

	basinSizes := FindBasinSizes(heightMap)
	sort.Slice(basinSizes, func(i, j int) bool { return basinSizes[i] > basinSizes[j] })
	mul := utils.Mul(basinSizes[0:3])
	log.Printf("sizeMul=%d\n", mul)
}

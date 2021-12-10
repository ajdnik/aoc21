package main

import (
	"log"
	"sort"

	"github.com/ajdnik/aoc21/day9"
	"github.com/ajdnik/aoc21/utils"
)

const NoBasin = -1

func MergeBasins(basins [][]int64, indexes []int64) ([][]int64, int64) {
	primary := indexes[0]
	remain := indexes[1:]
	for row := 0; row < len(basins); row++ {
		for col := 0; col < len(basins[row]); col++ {
			if utils.IsIncluded64(remain, basins[row][col]) {
				basins[row][col] = primary
			}
		}
	}
	return basins, primary
}

func FindBasinSizes(heights [][]int64) []int64 {
	basins := make([][]int64, len(heights))
	for row := 0; row < len(heights); row++ {
		basins[row] = make([]int64, len(heights[row]))
		for col := 0; col < len(heights[row]); col++ {
			basins[row][col] = NoBasin
		}
	}
	var count int64
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[row]); col++ {
			if heights[row][col] == 9 {
				continue
			}
			neighbours := day9.FindNeighbors(basins, row, col, func(itm int64) bool {
				return itm != NoBasin
			})
			if len(neighbours) == 0 {
				basins[row][col] = count
				count++
			} else if len(neighbours) == 1 {
				basins[row][col] = neighbours[0]
			} else {
				basins, idx := MergeBasins(basins, neighbours)
				basins[row][col] = idx
			}
		}
	}
	sizes := map[int64]int64{}
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[row]); col++ {
			if basins[row][col] == NoBasin {
				continue
			}
			if _, ok := sizes[basins[row][col]]; !ok {
				sizes[basins[row][col]] = 0
			}
			sizes[basins[row][col]]++
		}
	}
	basinSizes := []int64{}
	for _, val := range sizes {
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

	heights := [][]int64{}
	for scanner.Scan() {
		data := scanner.Text()
		res, err := day9.ToHeights(data)
		if err != nil {
			log.Fatal(err)
		}
		heights = append(heights, res)
	}

	basinSizes := FindBasinSizes(heights)
	sort.Slice(basinSizes, func(i, j int) bool { return basinSizes[i] > basinSizes[j] })
	mul := utils.Mul(basinSizes[0:3])
	log.Printf("sizeMul=%d\n", mul)
}

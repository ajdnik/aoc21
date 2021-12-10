package main

import (
	"log"

	"github.com/ajdnik/aoc21/day9"
	"github.com/ajdnik/aoc21/utils"
)

func FindLowPoints(heights [][]int64) []int64 {
	lowPoints := []int64{}
	for row := 0; row < len(heights); row++ {
		for col := 0; col < len(heights[row]); col++ {
			res := day9.FindNeighbors(heights, row, col, func(itm int64) bool {
				return itm <= heights[row][col]
			})
			if len(res) == 0 {
				lowPoints = append(lowPoints, heights[row][col])
			}
		}
	}
	return lowPoints
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

	lowPoints := FindLowPoints(heights)
	risks := utils.Add(lowPoints, 1)
	sum := utils.Sum(risks)
	log.Printf("riskSum=%d\n", sum)
}

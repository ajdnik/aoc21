package main

import (
	"log"

	"github.com/ajdnik/aoc21/day9"
	"github.com/ajdnik/aoc21/utils"
)

func FindLowPoints(heights [][]int64) []int64 {
	lowPoints := []int64{}
	for rowIdx, row := range heights {
		for colIdx, curHeight := range row {
			isLow := true
			if colIdx-1 >= 0 && row[colIdx-1] <= curHeight {
				isLow = false
			}
			if colIdx+1 < len(row) && row[colIdx+1] <= curHeight {
				isLow = false
			}
			if rowIdx-1 >= 0 && heights[rowIdx-1][colIdx] <= curHeight {
				isLow = false
			}
			if rowIdx+1 < len(heights) && heights[rowIdx+1][colIdx] <= curHeight {
				isLow = false
			}
			if isLow {
				lowPoints = append(lowPoints, curHeight)
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

	heightMap := [][]int64{}
	for scanner.Scan() {
		data := scanner.Text()
		heights, err := day9.ToHeights(data)
		if err != nil {
			log.Fatal(err)
		}
		heightMap = append(heightMap, heights)
	}

	lowPoints := FindLowPoints(heightMap)
	risks := []int64{}
	for _, height := range lowPoints {
		risks = append(risks, height+1)
	}
	sum := utils.Sum(risks)
	log.Printf("riskSum=%d\n", sum)
}

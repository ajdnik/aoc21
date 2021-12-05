package main

import (
	"log"

	"github.com/ajdnik/aoc21/day5"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	lines := []day5.Line{}
	var max int64
	for scanner.Scan() {
		data := scanner.Text()
		line, err := day5.ToLine(data)
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, line)
		if max < line.MaxDimentsion() {
			max = line.MaxDimentsion()
		}
	}

	field := make([]int64, (max+1)*(max+1))
	for _, line := range lines {
		if !line.IsHorizOrVert() {
			continue
		}
		for next, point, ok := line.PointGenerator(), (day5.Point{}), true; ok; {
			point, ok = next()
			field[point.X*max+point.Y]++
		}
	}

	var overlapping int64
	for _, num := range field {
		if num >= 2 {
			overlapping++
		}
	}

	log.Printf("overlapping=%d\n", overlapping)
}

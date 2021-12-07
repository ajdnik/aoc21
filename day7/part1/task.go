package main

import (
	"log"
	"math"

	"github.com/ajdnik/aoc21/day7"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	var positions []int64
	for scanner.Scan() {
		data := scanner.Text()
		positions, err = utils.ToIntList(data, ",")
		if err != nil {
			log.Fatal(err)
		}
	}

	minFuel := day7.FindMinFuel(positions, func(cur, goal int64) int64 {
		return int64(math.Abs(float64(cur - goal)))
	})

	log.Printf("minFuel=%d\n", minFuel)
}

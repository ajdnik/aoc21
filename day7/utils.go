package day7

import (
	"math"

	"github.com/ajdnik/aoc21/utils"
)

func FindMinFuel(positions []int64, getFuel func(cur, goal int64) int64) int64 {
	min, max := utils.MinMax(positions)

	minFuel := int64(math.MaxInt64)
	for goal := min; goal < max; goal++ {
		var fuel int64
		for _, pos := range positions {
			fuel += getFuel(pos, goal)
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

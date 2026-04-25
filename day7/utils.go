package day7

import (
	"math"
	"slices"
)

func FindMinFuel(positions []int64, getFuel func(cur, goal int64) int64) int64 {
	min := slices.Min(positions)
	max := slices.Max(positions)

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

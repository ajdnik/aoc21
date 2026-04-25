package day07

import (
	"math"
	"slices"

	"github.com/ajdnik/aoc21/utils"
)

func findMinFuel(positions []int64, getFuel func(cur, goal int64) int64) int64 {
	mn := slices.Min(positions)
	mx := slices.Max(positions)

	minFuel := int64(math.MaxInt64)
	for goal := mn; goal < mx; goal++ {
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

func Part1(lines []string) int64 {
	positions, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}
	return findMinFuel(positions, func(cur, goal int64) int64 {
		return int64(math.Abs(float64(cur - goal)))
	})
}

func Part2(lines []string) int64 {
	positions, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}
	return findMinFuel(positions, func(cur, goal int64) int64 {
		return utils.SumUpTo(int64(math.Abs(float64(cur - goal))))
	})
}

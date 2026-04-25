package day07

import (
	"math"
	"slices"

	"github.com/ajdnik/aoc21/utils"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sumUpTo(n int) int {
	return n * (n + 1) / 2
}

func findMinFuel(positions []int, getFuel func(cur, goal int) int) int {
	mn := slices.Min(positions)
	mx := slices.Max(positions)

	minFuel := math.MaxInt
	for goal := mn; goal < mx; goal++ {
		var fuel int
		for _, pos := range positions {
			fuel += getFuel(pos, goal)
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func Part1(lines []string) int {
	positions, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}
	return findMinFuel(positions, func(cur, goal int) int {
		return abs(cur - goal)
	})
}

func Part2(lines []string) int {
	positions, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}
	return findMinFuel(positions, func(cur, goal int) int {
		return sumUpTo(abs(cur - goal))
	})
}

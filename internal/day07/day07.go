// Package day07 solves AoC 2021 day 7: The Treachery of Whales.
// Find the optimal alignment position for crab submarines to minimize fuel.
package day07

import (
	"math"
	"slices"

	"github.com/ajdnik/aoc21/utils"
)

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

// Part1 uses constant fuel cost per step.
func Part1(lines []string) int {
	positions, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}
	return findMinFuel(positions, func(cur, goal int) int {
		return utils.Abs(cur - goal)
	})
}

// Part2 uses triangular fuel cost (1+2+3+...+n per n steps).
func Part2(lines []string) int {
	positions, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}
	return findMinFuel(positions, func(cur, goal int) int {
		return sumUpTo(utils.Abs(cur - goal))
	})
}

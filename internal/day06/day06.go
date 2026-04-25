// Package day06 solves AoC 2021 day 6: Lanternfish.
// Simulate lanternfish population growth using bucket counting for efficiency.
package day06

import (
	"github.com/ajdnik/aoc21/utils"
)

// simulate counts fish after the given number of days.
// Each bucket[i] holds the count of fish with timer value i (0-8).
// Fish at 0 spawn a new fish at 8 and reset to 6.
func simulate(timers []int, days int) int {
	var buckets [9]int
	for _, t := range timers {
		buckets[t]++
	}
	for day := 0; day < days; day++ {
		var next [9]int
		next[6] = buckets[0]
		next[8] = buckets[0]
		for i := 1; i <= 8; i++ {
			next[i-1] += buckets[i]
		}
		buckets = next
	}
	total := 0
	for _, count := range buckets {
		total += count
	}
	return total
}

func Part1(lines []string) int {
	timers, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}
	return simulate(timers, 80)
}

func Part2(lines []string) int {
	timers, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}
	return simulate(timers, 256)
}

package day06

import (
	"github.com/ajdnik/aoc21/utils"
)

func simulateDay(timers []int) []int {
	for idx, timer := range timers {
		if timer == 0 {
			timers = append(timers, 8)
			timers[idx] = 6
		} else {
			timers[idx]--
		}
	}
	return timers
}

func simulateDays(timers []int, days int) []int {
	for day := 0; day < days; day++ {
		timers = simulateDay(timers)
	}
	return timers
}

func Part1(lines []string) int {
	timers, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}
	timers = simulateDays(timers, 80)
	return len(timers)
}

func Part2(lines []string) int {
	initialTimers, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}

	cache := map[int]int{}
	var totalFishes int
	for _, initialTimer := range initialTimers {
		if val, ok := cache[initialTimer]; ok {
			totalFishes += val
			continue
		}
		timers := []int{initialTimer}
		timers = simulateDays(timers, 256)
		cache[initialTimer] = len(timers)
		totalFishes += len(timers)
	}
	return totalFishes
}

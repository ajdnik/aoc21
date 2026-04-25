package day06

import (
	"github.com/ajdnik/aoc21/utils"
)

func simulateDay(timers []int64) []int64 {
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

func simulateDays(timers []int64, days int) []int64 {
	for day := 0; day < days; day++ {
		timers = simulateDay(timers)
	}
	return timers
}

func Part1(lines []string) int64 {
	timers, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}
	timers = simulateDays(timers, 80)
	return int64(len(timers))
}

func Part2(lines []string) uint64 {
	initialTimers, err := utils.ToIntList(lines[0], ",")
	if err != nil {
		panic(err)
	}

	cache := map[int64]uint64{}
	var totalFishes uint64
	for _, initialTimer := range initialTimers {
		if val, ok := cache[initialTimer]; ok {
			totalFishes += val
			continue
		}
		timers := []int64{initialTimer}
		timers = simulateDays(timers, 256)
		popLen := uint64(len(timers))
		cache[initialTimer] = popLen
		totalFishes += popLen
	}
	return totalFishes
}

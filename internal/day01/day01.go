// Package day01 solves AoC 2021 day 1: Sonar Sweep.
// Count depth measurement increases, both individually and in sliding windows of 3.
package day01

import "strconv"

func parseNumbers(lines []string) []int {
	nums := make([]int, len(lines))
	for i, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

// Part1 counts how many measurements are larger than the previous one.
func Part1(lines []string) int {
	nums := parseNumbers(lines)
	var inc int
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			inc++
		}
	}
	return inc
}

// Part2 counts increases using a sliding window of 3 measurements.
func Part2(lines []string) int {
	nums := parseNumbers(lines)
	var inc int
	prev := nums[0] + nums[1] + nums[2]
	for i := 3; i < len(nums); i++ {
		sum := nums[i] + nums[i-1] + nums[i-2]
		if sum > prev {
			inc++
		}
		prev = sum
	}
	return inc
}

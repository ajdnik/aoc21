package day01

import (
	"github.com/ajdnik/aoc21/utils"
)

func parseNumbers(lines []string) []int64 {
	nums := make([]int64, len(lines))
	for i, line := range lines {
		num, err := utils.ToInt(line)
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}

func Part1(lines []string) int64 {
	nums := parseNumbers(lines)
	var inc int64
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			inc++
		}
	}
	return inc
}

func Part2(lines []string) int64 {
	nums := parseNumbers(lines)
	var inc int64
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

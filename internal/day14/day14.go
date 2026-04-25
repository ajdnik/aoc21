// Package day14 solves AoC 2021 day 14: Extended Polymerization.
// Apply pair insertion rules to a polymer, tracking pair counts for efficiency.
package day14

import (
	"math"
	"strings"
)

func parseInput(lines []string) (string, map[string]byte) {
	template := lines[0]
	rules := make(map[string]byte)
	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1][0]
	}
	return template, rules
}

func solve(lines []string, steps int) int {
	template, rules := parseInput(lines)

	pairs := map[string]int{}
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for range steps {
		next := map[string]int{}
		for pair, count := range pairs {
			if insert, ok := rules[pair]; ok {
				next[string([]byte{pair[0], insert})] += count
				next[string([]byte{insert, pair[1]})] += count
			} else {
				next[pair] += count
			}
		}
		pairs = next
	}

	counts := map[byte]int{}
	for pair, count := range pairs {
		counts[pair[0]] += count
	}
	counts[template[len(template)-1]]++

	minCount := math.MaxInt
	maxCount := 0
	for _, count := range counts {
		if count < minCount {
			minCount = count
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount - minCount
}

func Part1(lines []string) int {
	return solve(lines, 10)
}

func Part2(lines []string) int {
	return solve(lines, 40)
}

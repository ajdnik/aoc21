// Package day10 solves AoC 2021 day 10: Syntax Scoring.
// Detect corrupted and incomplete bracket sequences.
package day10

import "slices"

const emptyRune = rune('E')

func parseLine(line string) ([]rune, rune) {
	var stack []rune
	for _, char := range line {
		if char == '(' || char == '[' || char == '{' || char == '<' {
			stack = append(stack, char)
			continue
		}
		if len(stack) == 0 {
			return nil, char
		}
		popped := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if popped == '(' && char != ')' {
			return stack, char
		}
		if popped == '[' && char != ']' {
			return stack, char
		}
		if popped == '{' && char != '}' {
			return stack, char
		}
		if popped == '<' && char != '>' {
			return stack, char
		}
	}
	return stack, emptyRune
}

var syntaxScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

// Part1 scores corrupted lines by their first illegal closing character.
func Part1(lines []string) int {
	var score int
	for _, line := range lines {
		_, char := parseLine(line)
		if val, ok := syntaxScores[char]; ok {
			score += val
		}
	}
	return score
}

var completionScores = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

// Part2 scores incomplete lines by the characters needed to complete them.
func Part2(lines []string) int {
	var scores []int
	for _, line := range lines {
		stack, char := parseLine(line)
		if char != emptyRune {
			continue
		}
		var score int
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			score *= 5
			score += completionScores[top]
		}
		scores = append(scores, score)
	}
	slices.SortFunc(scores, func(a, b int) int { return b - a })
	return scores[len(scores)/2]
}

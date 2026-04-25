package day08

import (
	"sort"
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

type observation struct {
	Patterns []string
	Output   []string
}

func toObservation(data string) *observation {
	data = utils.NormalizeSpaces(data)
	split := strings.Split(data, "|")
	patterns := strings.Split(split[0], " ")
	output := strings.Split(split[1], " ")
	return &observation{
		Patterns: patterns,
		Output:   output,
	}
}

func parseObservations(lines []string) []*observation {
	obs := make([]*observation, len(lines))
	for i, line := range lines {
		obs[i] = toObservation(line)
	}
	return obs
}

func Part1(lines []string) int64 {
	observations := parseObservations(lines)
	var sum int64
	for _, observ := range observations {
		for _, out := range observ.Output {
			switch len(out) {
			case 2, 3, 4, 7:
				sum++
			}
		}
	}
	return sum
}

func getUnknownChars(input string, known string) string {
	unknown := input
	for _, char := range known {
		unknown = strings.Replace(unknown, string(char), "", -1)
	}
	return unknown
}

func findCharsNotInAll(inputs []string) string {
	count := map[rune]int{}
	for _, in := range inputs {
		for _, char := range in {
			count[char]++
		}
	}
	var unique string
	for key, val := range count {
		if val != len(inputs) {
			unique += string(key)
		}
	}
	return unique
}

func findCharsInBoth(input1, input2 string) string {
	chars := map[rune]bool{}
	for _, char := range input1 {
		chars[char] = true
	}
	var overlapping string
	for _, char := range input2 {
		if chars[char] {
			overlapping += string(char)
		}
	}
	return overlapping
}

func solveWires(patterns []string) map[string]string {
	byLen := map[int][]string{}
	for _, pat := range patterns {
		l := len(pat)
		byLen[l] = append(byLen[l], pat)
	}

	segA := findCharsNotInAll([]string{byLen[3][0], byLen[2][0]})
	segCorF := getUnknownChars(byLen[3][0], segA)
	segBorE := getUnknownChars(findCharsNotInAll(byLen[5]), segCorF)
	segDorCorE := findCharsNotInAll(byLen[6])
	segC := findCharsInBoth(segCorF, segDorCorE)
	segE := findCharsInBoth(segBorE, segDorCorE)
	segF := getUnknownChars(segCorF, segC)
	segB := getUnknownChars(segBorE, segE)
	segD := getUnknownChars(segDorCorE, segC+segE)
	segG := getUnknownChars(byLen[7][0], segA+segB+segC+segD+segE+segF)

	return map[string]string{
		segA: "a",
		segB: "b",
		segC: "c",
		segD: "d",
		segE: "e",
		segF: "f",
		segG: "g",
	}
}

func toCorrectSegments(output string, mapping map[string]string) string {
	var seg []string
	for _, char := range output {
		if val, ok := mapping[string(char)]; ok {
			seg = append(seg, val)
		}
	}
	sort.Strings(seg)
	return strings.Join(seg, "")
}

func toNumbers(output []string, mapping map[string]string) []int {
	charToNum := map[string]int{
		"cf":      1,
		"acf":     7,
		"bcdf":    4,
		"acdeg":   2,
		"acdfg":   3,
		"abdfg":   5,
		"abcefg":  0,
		"abdefg":  6,
		"abcdfg":  9,
		"abcdefg": 8,
	}
	var numbers []int
	for _, out := range output {
		seg := toCorrectSegments(out, mapping)
		if val, ok := charToNum[seg]; ok {
			numbers = append(numbers, val)
		}
	}
	return numbers
}

func Part2(lines []string) int64 {
	observations := parseObservations(lines)
	var sum int64
	for _, observ := range observations {
		mapping := solveWires(observ.Patterns)
		numbers := toNumbers(observ.Output, mapping)
		sum += int64(numbers[0]*1000 + numbers[1]*100 + numbers[2]*10 + numbers[3])
	}
	return sum
}

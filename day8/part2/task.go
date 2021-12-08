package main

import (
	"log"
	"sort"
	"strings"

	"github.com/ajdnik/aoc21/day8"
	"github.com/ajdnik/aoc21/utils"
)

func GetUnknownChars(input string, known string) string {
	unknown := input
	for _, char := range known {
		unknown = strings.Replace(unknown, string(char), "", -1)
	}
	return unknown
}

func FindCharsNotInAll(inputs []string) string {
	count := map[rune]int{}
	for _, in := range inputs {
		for _, char := range in {
			if _, ok := count[char]; !ok {
				count[char] = 1
			} else {
				count[char]++
			}
		}
	}
	unique := ""
	for key, val := range count {
		if val != len(inputs) {
			unique += string(key)
		}
	}
	return unique
}

func FindCharsInBoth(input1, input2 string) string {
	chars := map[rune]bool{}
	for _, char := range input1 {
		if _, ok := chars[char]; !ok {
			chars[char] = true
		}
	}
	overlapping := ""
	for _, char := range input2 {
		if _, ok := chars[char]; ok {
			overlapping += string(char)
		}
	}
	return overlapping
}

func SolveWires(patterns []string) map[string]string {
	byLen := map[int][]string{}
	for _, pat := range patterns {
		l := len(pat)
		if _, ok := byLen[l]; !ok {
			byLen[l] = []string{pat}
		} else {
			byLen[l] = append(byLen[l], pat)
		}
	}

	segA := FindCharsNotInAll([]string{byLen[3][0], byLen[2][0]})
	segCorF := GetUnknownChars(byLen[3][0], segA)
	segBorE := GetUnknownChars(FindCharsNotInAll(byLen[5]), segCorF)
	segDorCorE := FindCharsNotInAll(byLen[6])
	segC := FindCharsInBoth(segCorF, segDorCorE)
	segE := FindCharsInBoth(segBorE, segDorCorE)
	segF := GetUnknownChars(segCorF, segC)
	segB := GetUnknownChars(segBorE, segE)
	segD := GetUnknownChars(segDorCorE, segC+segE)
	segG := GetUnknownChars(byLen[7][0], segA+segB+segC+segD+segE+segF)

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

func ToCorrectSegments(output string, mapping map[string]string) string {
	seg := []string{}
	for _, char := range output {
		if val, ok := mapping[string(char)]; ok {
			seg = append(seg, val)
		}
	}
	sort.Strings(seg)
	return strings.Join(seg, "")
}

func ToNumbers(output []string, mapping map[string]string) []int {
	numbers := []int{}
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
	for _, out := range output {
		seg := ToCorrectSegments(out, mapping)
		if val, ok := charToNum[seg]; ok {
			numbers = append(numbers, val)
		}
	}
	return numbers
}

func ToNumber(observation *day8.Observation) uint64 {
	mapping := SolveWires(observation.Patterns)
	numbers := ToNumbers(observation.Output, mapping)
	return uint64(numbers[0]*1000 + numbers[1]*100 + numbers[2]*10 + numbers[3])
}

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	observations := []*day8.Observation{}
	for scanner.Scan() {
		data := scanner.Text()
		observation := day8.ToObservation(data)
		observations = append(observations, observation)
	}

	var sum uint64
	for _, observ := range observations {
		sum += ToNumber(observ)
	}
	log.Printf("sum=%d\n", sum)
}

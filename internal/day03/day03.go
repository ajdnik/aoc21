// Package day03 solves AoC 2021 day 3: Binary Diagnostic.
// Analyze binary diagnostic reports to compute power and life support ratings.
package day03

import (
	"strconv"
	"strings"
)

type searchType int

const (
	oxygen searchType = iota
	co2
)

func binaryToInt(data string) int {
	i, err := strconv.ParseInt(data, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func filterMatches(data []string, match byte, pos int) []string {
	var filtered []string
	for _, itm := range data {
		if itm[pos] == match {
			filtered = append(filtered, itm)
		}
	}
	return filtered
}

func buildValue(data []string, typ searchType, pos int) string {
	var zeros, ones int
	for _, itm := range data {
		switch itm[pos] {
		case '1':
			ones++
		case '0':
			zeros++
		}
	}

	var filtered []string
	switch typ {
	case oxygen:
		if ones >= zeros {
			filtered = filterMatches(data, '1', pos)
		} else {
			filtered = filterMatches(data, '0', pos)
		}
	case co2:
		if ones < zeros {
			filtered = filterMatches(data, '1', pos)
		} else {
			filtered = filterMatches(data, '0', pos)
		}
	}

	if len(filtered) == 1 {
		return filtered[0]
	}
	return buildValue(filtered, typ, pos+1)
}

// Part1 computes power consumption: gamma rate * epsilon rate.
func Part1(lines []string) int {
	var ones, zeros []int
	for _, data := range lines {
		if ones == nil {
			ones = make([]int, len(data))
			zeros = make([]int, len(data))
		}
		for idx, char := range data {
			switch char {
			case '1':
				ones[idx]++
			case '0':
				zeros[idx]++
			}
		}
	}

	var gamma, epsilon strings.Builder
	for i, z := range zeros {
		if z > ones[i] {
			gamma.WriteByte('0')
			epsilon.WriteByte('1')
		} else {
			gamma.WriteByte('1')
			epsilon.WriteByte('0')
		}
	}

	return binaryToInt(gamma.String()) * binaryToInt(epsilon.String())
}

// Part2 computes life support rating: oxygen generator * CO2 scrubber values.
func Part2(lines []string) int {
	oxygenVal := buildValue(lines, oxygen, 0)
	co2Val := buildValue(lines, co2, 0)
	return binaryToInt(oxygenVal) * binaryToInt(co2Val)
}

package day03

import (
	"github.com/ajdnik/aoc21/utils"
)

type searchType int

const (
	oxygen searchType = iota
	co2
)

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
	var zeros, ones uint64
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

func Part1(lines []string) int64 {
	var ones, zeros []uint64
	for _, data := range lines {
		if ones == nil {
			ones = make([]uint64, len(data))
			zeros = make([]uint64, len(data))
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

	var gamma, epsilon string
	for i := 0; i < len(zeros); i++ {
		if zeros[i] > ones[i] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaInt, err := utils.BinaryToInt(gamma)
	if err != nil {
		panic(err)
	}
	epsilonInt, err := utils.BinaryToInt(epsilon)
	if err != nil {
		panic(err)
	}
	return gammaInt * epsilonInt
}

func Part2(lines []string) int64 {
	oxygenVal := buildValue(lines, oxygen, 0)
	co2Val := buildValue(lines, co2, 0)

	oxygenInt, err := utils.BinaryToInt(oxygenVal)
	if err != nil {
		panic(err)
	}
	co2Int, err := utils.BinaryToInt(co2Val)
	if err != nil {
		panic(err)
	}
	return oxygenInt * co2Int
}

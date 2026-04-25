package day03

import "strconv"

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

	return binaryToInt(gamma) * binaryToInt(epsilon)
}

func Part2(lines []string) int {
	oxygenVal := buildValue(lines, oxygen, 0)
	co2Val := buildValue(lines, co2, 0)
	return binaryToInt(oxygenVal) * binaryToInt(co2Val)
}

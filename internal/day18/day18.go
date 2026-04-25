// Package day18 solves AoC 2021 day 18: Snailfish.
// Add and reduce nested pairs using string manipulation for explode/split rules.
package day18

import (
	"strconv"
)

func addNum(a, b string) string {
	return reduce("[" + a + "," + b + "]")
}

// reduce repeatedly applies explode (priority) then split until stable.
func reduce(s string) string {
	for {
		if next, changed := explode(s); changed {
			s = next
			continue
		}
		if next, changed := split(s); changed {
			s = next
			continue
		}
		break
	}
	return s
}

func explode(s string) (string, bool) {
	depth := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '[':
			depth++
			if depth > 4 {
				return doExplode(s, i), true
			}
		case ']':
			depth--
		}
	}
	return s, false
}

func doExplode(s string, pos int) string {
	// Find the pair starting at pos: [left,right]
	comma := pos + 1
	for s[comma] != ',' {
		comma++
	}
	close := comma + 1
	for s[close] != ']' {
		close++
	}
	left, _ := strconv.Atoi(s[pos+1 : comma])
	right, _ := strconv.Atoi(s[comma+1 : close])

	// Build result: add left to nearest left number, replace pair with 0, add right to nearest right number
	prefix := s[:pos]
	suffix := s[close+1:]

	// Add right value to first number on the right
	suffix = addToFirst(suffix, right)

	// Add left value to first number on the left (searching backwards)
	prefix = addToLast(prefix, left)

	return prefix + "0" + suffix
}

func addToFirst(s string, val int) string {
	// Find first number in s and add val to it
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			j := i
			for j < len(s) && isDigit(s[j]) {
				j++
			}
			n, _ := strconv.Atoi(s[i:j])
			return s[:i] + strconv.Itoa(n+val) + s[j:]
		}
	}
	return s
}

func addToLast(s string, val int) string {
	// Find last number in s and add val to it
	for i := len(s) - 1; i >= 0; i-- {
		if isDigit(s[i]) {
			j := i + 1
			for i > 0 && isDigit(s[i-1]) {
				i--
			}
			n, _ := strconv.Atoi(s[i:j])
			return s[:i] + strconv.Itoa(n+val) + s[j:]
		}
	}
	return s
}

func split(s string) (string, bool) {
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			j := i
			for j < len(s) && isDigit(s[j]) {
				j++
			}
			if j-i >= 2 {
				n, _ := strconv.Atoi(s[i:j])
				left := n / 2
				right := (n + 1) / 2
				replacement := "[" + strconv.Itoa(left) + "," + strconv.Itoa(right) + "]"
				return s[:i] + replacement + s[j:], true
			}
		}
	}
	return s, false
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

// magnitude computes 3*left + 2*right recursively.
func magnitude(s string) int {
	val, _ := parseMagnitude(s, 0)
	return val
}

func parseMagnitude(s string, pos int) (int, int) {
	if s[pos] == '[' {
		left, p := parseMagnitude(s, pos+1)
		// skip comma
		right, p := parseMagnitude(s, p+1)
		// skip ]
		return 3*left + 2*right, p + 1
	}
	j := pos
	for j < len(s) && isDigit(s[j]) {
		j++
	}
	n, _ := strconv.Atoi(s[pos:j])
	return n, j
}

func Part1(lines []string) int {
	result := lines[0]
	for i := 1; i < len(lines); i++ {
		result = addNum(result, lines[i])
	}
	return magnitude(result)
}

func Part2(lines []string) int {
	best := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}
			if m := magnitude(addNum(lines[i], lines[j])); m > best {
				best = m
			}
		}
	}
	return best
}

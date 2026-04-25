package day03

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := int64(198)
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := int64(230)
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

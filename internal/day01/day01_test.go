package day01

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`199
200
208
210
200
207
240
269
260
263`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 7
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := 5
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

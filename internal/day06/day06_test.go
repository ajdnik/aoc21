package day06

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`3,4,3,1,2`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 5934
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
got := Part2(testInput)
	want := 26984457539
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

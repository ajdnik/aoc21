package day09

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`2199943210
3987894921
9856789892
8767896789
9899965678`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 15
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := 1134
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

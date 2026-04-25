package day11

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 1656
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := 195
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

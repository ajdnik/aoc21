package day15

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 40
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := 315
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

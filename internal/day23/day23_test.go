package day23

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 12521
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := 44169
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

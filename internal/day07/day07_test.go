package day07

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`16,1,2,0,4,2,7,1,2,14`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 37
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := 168
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

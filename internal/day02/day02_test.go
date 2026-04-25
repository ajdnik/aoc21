package day02

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`forward 5
down 5
forward 8
up 3
down 8
forward 2`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := int64(150)
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := int64(900)
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

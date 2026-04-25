package day25

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 58
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

package day13

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 17
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := "#####\n#...#\n#...#\n#...#\n#####\n.....\n....."
	if got != want {
		t.Errorf("Part2() =\n%s\nwant:\n%s", got, want)
	}
}

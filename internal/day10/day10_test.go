package day10

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var testInput = utils.ReadLines(strings.NewReader(`[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`))

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 26397
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := 288957
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

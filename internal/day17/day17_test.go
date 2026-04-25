package day17

import (
	"testing"
)

var testInput = []string{"target area: x=20..30, y=-10..-5"}

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 45
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := 112
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

package day21

import (
	"testing"
)

var testInput = []string{
	"Player 1 starting position: 4",
	"Player 2 starting position: 8",
}

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	want := 739785
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	want := 444356092776315
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

package day03

import (
	"os"
	"strings"
	"testing"
)

func readTestInput(t *testing.T) []string {
	t.Helper()
	data, err := os.ReadFile("../../input/day03_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func TestPart1(t *testing.T) {
	lines := readTestInput(t)
	got := Part1(lines)
	want := int64(198)
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines := readTestInput(t)
	got := Part2(lines)
	want := int64(230)
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

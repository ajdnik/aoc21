package day06

import (
	"os"
	"strings"
	"testing"
)

func readTestInput(t *testing.T) []string {
	t.Helper()
	data, err := os.ReadFile("../../input/day06_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func TestPart1(t *testing.T) {
	lines := readTestInput(t)
	got := Part1(lines)
	want := int64(5934)
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping slow brute-force simulation in short mode")
	}
	lines := readTestInput(t)
	got := Part2(lines)
	want := uint64(26984457539)
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}

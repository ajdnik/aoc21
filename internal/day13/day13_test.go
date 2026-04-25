package day13

import (
	"os"
	"strings"
	"testing"
)

func readTestInput(t *testing.T) []string {
	t.Helper()
	data, err := os.ReadFile("../../input/day13_test.txt")
	if err != nil {
		t.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func TestPart1(t *testing.T) {
	lines := readTestInput(t)
	got := Part1(lines)
	want := int64(17)
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines := readTestInput(t)
	got := Part2(lines)
	want := "#####\n#...#\n#...#\n#...#\n#####\n.....\n....."
	if got != want {
		t.Errorf("Part2() =\n%s\nwant:\n%s", got, want)
	}
}

package day12

import (
	"os"
	"strings"
	"testing"
)

func readTestInput(t *testing.T, name string) []string {
	t.Helper()
	data, err := os.ReadFile("../../input/" + name)
	if err != nil {
		t.Fatal(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func TestPart1(t *testing.T) {
	tests := []struct {
		file string
		want int64
	}{
		{"day12_test.txt", 10},
		{"day12_test2.txt", 19},
		{"day12_test3.txt", 226},
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			lines := readTestInput(t, tt.file)
			got := Part1(lines)
			if got != tt.want {
				t.Errorf("Part1() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		file string
		want int64
	}{
		{"day12_test.txt", 36},
		{"day12_test2.txt", 103},
		{"day12_test3.txt", 3509},
	}
	for _, tt := range tests {
		t.Run(tt.file, func(t *testing.T) {
			lines := readTestInput(t, tt.file)
			got := Part2(lines)
			if got != tt.want {
				t.Errorf("Part2() = %d, want %d", got, tt.want)
			}
		})
	}
}

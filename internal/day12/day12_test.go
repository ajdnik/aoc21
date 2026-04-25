package day12

import (
	"strings"
	"testing"

	"github.com/ajdnik/aoc21/utils"
)

var (
	testInput1 = utils.ReadLines(strings.NewReader(`start-A
start-b
A-c
A-b
b-d
A-end
b-end`))

	testInput2 = utils.ReadLines(strings.NewReader(`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`))

	testInput3 = utils.ReadLines(strings.NewReader(`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`))
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"small", testInput1, 10},
		{"medium", testInput2, 19},
		{"large", testInput3, 226},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)
			if got != tt.want {
				t.Errorf("Part1() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{"small", testInput1, 36},
		{"medium", testInput2, 103},
		{"large", testInput3, 3509},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.input)
			if got != tt.want {
				t.Errorf("Part2() = %d, want %d", got, tt.want)
			}
		})
	}
}

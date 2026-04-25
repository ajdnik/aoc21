package day02

import (
	"errors"
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

type direction int

const (
	forward direction = iota
	up
	down
)

type movement struct {
	dir  direction
	unit int
}

func toMovement(data string) (movement, error) {
	var mov movement
	data = utils.NormalizeSpaces(data)
	parts := strings.Split(data, " ")
	if len(parts) != 2 {
		return mov, errors.New("movement input needs two elements")
	}

	switch strings.ToLower(parts[0]) {
	case "forward":
		mov.dir = forward
	case "up":
		mov.dir = up
	case "down":
		mov.dir = down
	default:
		return mov, errors.New("unknown movement direction")
	}

	unit, err := utils.ToInt(parts[1])
	if err != nil {
		return mov, err
	}
	mov.unit = unit
	return mov, nil
}

func parseMovements(lines []string) []movement {
	movs := make([]movement, len(lines))
	for i, line := range lines {
		mov, err := toMovement(line)
		if err != nil {
			panic(err)
		}
		movs[i] = mov
	}
	return movs
}

func Part1(lines []string) int {
	movs := parseMovements(lines)
	var horiz, depth int
	for _, mov := range movs {
		switch mov.dir {
		case forward:
			horiz += mov.unit
		case up:
			depth -= mov.unit
		case down:
			depth += mov.unit
		}
	}
	return horiz * depth
}

func Part2(lines []string) int {
	movs := parseMovements(lines)
	var horiz, depth, aim int
	for _, mov := range movs {
		switch mov.dir {
		case forward:
			horiz += mov.unit
			depth += aim * mov.unit
		case up:
			aim -= mov.unit
		case down:
			aim += mov.unit
		}
	}
	return horiz * depth
}

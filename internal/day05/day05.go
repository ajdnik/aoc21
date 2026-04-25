package day05

import (
	"errors"
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

type direction int

const (
	ascending direction = iota
	descending
	static
)

type point struct {
	X int64
	Y int64
}

type line struct {
	Start point
	End   point
}

func (l line) maxDimension() int64 {
	m := l.Start.X
	if l.Start.Y > m {
		m = l.Start.Y
	}
	if l.End.X > m {
		m = l.End.X
	}
	if l.End.Y > m {
		m = l.End.Y
	}
	return m
}

func (l line) isHorizOrVert() bool {
	return l.Start.X == l.End.X || l.Start.Y == l.End.Y
}

func (l line) pointGenerator() func() (point, bool) {
	curPoint, xDir, yDir := l.initializePointGenerator()
	return func() (point, bool) {
		switch xDir {
		case ascending:
			curPoint.X++
		case descending:
			curPoint.X--
		}
		switch yDir {
		case ascending:
			curPoint.Y++
		case descending:
			curPoint.Y--
		}
		hasNext := !(curPoint.X == l.End.X && curPoint.Y == l.End.Y)
		return curPoint, hasNext
	}
}

func (l line) initializePointGenerator() (point, direction, direction) {
	pt := point{}
	var xDir, yDir direction
	if l.Start.X == l.End.X {
		xDir = static
		pt.X = l.Start.X
	} else if l.Start.X < l.End.X {
		xDir = ascending
		pt.X = l.Start.X - 1
	} else {
		xDir = descending
		pt.X = l.Start.X + 1
	}
	if l.Start.Y == l.End.Y {
		yDir = static
		pt.Y = l.Start.Y
	} else if l.Start.Y < l.End.Y {
		yDir = ascending
		pt.Y = l.Start.Y - 1
	} else {
		yDir = descending
		pt.Y = l.Start.Y + 1
	}
	return pt, xDir, yDir
}

func toLine(data string) (line, error) {
	var l line
	data = utils.NormalizeSpaces(data)
	firstSplit := strings.Split(data, ",")
	if len(firstSplit) != 3 {
		return l, errors.New("parsing failed: more than 2 commas found")
	}
	x1, err := utils.ToInt(firstSplit[0])
	if err != nil {
		return l, err
	}
	l.Start.X = x1
	y2, err := utils.ToInt(firstSplit[2])
	if err != nil {
		return l, err
	}
	l.End.Y = y2
	secondSplit := strings.Split(firstSplit[1], " ")
	if len(secondSplit) != 3 {
		return l, errors.New("parsing failed: more than 2 spaces found")
	}
	y1, err := utils.ToInt(secondSplit[0])
	if err != nil {
		return l, err
	}
	l.Start.Y = y1
	x2, err := utils.ToInt(secondSplit[2])
	if err != nil {
		return l, err
	}
	l.End.X = x2
	return l, nil
}

func parseLines(data []string) ([]line, int64) {
	var lines []line
	var max int64
	for _, d := range data {
		l, err := toLine(d)
		if err != nil {
			panic(err)
		}
		lines = append(lines, l)
		if max < l.maxDimension() {
			max = l.maxDimension()
		}
	}
	return lines, max
}

func countOverlapping(field []int64) int64 {
	var count int64
	for _, num := range field {
		if num >= 2 {
			count++
		}
	}
	return count
}

func Part1(data []string) int64 {
	lines, max := parseLines(data)

	field := make([]int64, (max+1)*(max+1))
	for _, l := range lines {
		if !l.isHorizOrVert() {
			continue
		}
		for next, pt, ok := l.pointGenerator(), (point{}), true; ok; {
			pt, ok = next()
			field[pt.X*max+pt.Y]++
		}
	}
	return countOverlapping(field)
}

func Part2(data []string) int64 {
	lines, max := parseLines(data)

	field := make([]int64, (max+1)*(max+1))
	for _, l := range lines {
		for next, pt, ok := l.pointGenerator(), (point{}), true; ok; {
			pt, ok = next()
			field[pt.X*max+pt.Y]++
		}
	}
	return countOverlapping(field)
}

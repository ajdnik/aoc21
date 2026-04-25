package day05

import "fmt"

type direction int

const (
	ascending direction = iota
	descending
	static
)

type point struct {
	X int
	Y int
}

type line struct {
	Start point
	End   point
}

func (l line) maxDimension() int {
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

func toLine(data string) line {
	var l line
	fmt.Sscanf(data, "%d,%d -> %d,%d", &l.Start.X, &l.Start.Y, &l.End.X, &l.End.Y)
	return l
}

func parseLines(data []string) ([]line, int) {
	var lines []line
	var max int
	for _, d := range data {
		l := toLine(d)
		lines = append(lines, l)
		if m := l.maxDimension(); max < m {
			max = m
		}
	}
	return lines, max
}

func countOverlapping(field []int) int {
	var count int
	for _, num := range field {
		if num >= 2 {
			count++
		}
	}
	return count
}

func Part1(data []string) int {
	lines, max := parseLines(data)

	field := make([]int, (max+1)*(max+1))
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

func Part2(data []string) int {
	lines, max := parseLines(data)

	field := make([]int, (max+1)*(max+1))
	for _, l := range lines {
		for next, pt, ok := l.pointGenerator(), (point{}), true; ok; {
			pt, ok = next()
			field[pt.X*max+pt.Y]++
		}
	}
	return countOverlapping(field)
}

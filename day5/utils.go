package day5

import (
	"errors"
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

type Direction int

const (
	Ascending Direction = iota
	Descending
	Static
)

type Point struct {
	X int64
	Y int64
}

type Line struct {
	Start Point
	End   Point
}

func (l Line) MaxDimentsion() int64 {
	max := l.Start.X
	if l.Start.Y > max {
		max = l.Start.Y
	}
	if l.End.X > max {
		max = l.End.X
	}
	if l.End.Y > max {
		max = l.End.Y
	}
	return max
}

func (l Line) IsHorizOrVert() bool {
	return l.Start.X == l.End.X || l.Start.Y == l.End.Y
}

func (l Line) PointGenerator() func() (Point, bool) {
	curPoint, xDir, yDir := l.initializePointGenerator()
	return func() (Point, bool) {
		switch xDir {
		case Ascending:
			curPoint.X++
		case Descending:
			curPoint.X--
		}
		switch yDir {
		case Ascending:
			curPoint.Y++
		case Descending:
			curPoint.Y--
		}
		hasNext := !(curPoint.X == l.End.X && curPoint.Y == l.End.Y)
		return curPoint, hasNext
	}
}

func (l Line) initializePointGenerator() (Point, Direction, Direction) {
	point := Point{}
	var xDir, yDir Direction
	if l.Start.X == l.End.X {
		xDir = Static
		point.X = l.Start.X
	} else if l.Start.X < l.End.X {
		xDir = Ascending
		point.X = l.Start.X - 1
	} else {
		xDir = Descending
		point.X = l.Start.X + 1
	}
	if l.Start.Y == l.End.Y {
		yDir = Static
		point.Y = l.Start.Y
	} else if l.Start.Y < l.End.Y {
		yDir = Ascending
		point.Y = l.Start.Y - 1
	} else {
		yDir = Descending
		point.Y = l.Start.Y + 1
	}
	return point, xDir, yDir
}

func ToLine(data string) (Line, error) {
	line := Line{}
	data = utils.NormalizeSpaces(data)
	firstSplit := strings.Split(data, ",")
	if len(firstSplit) != 3 {
		return line, errors.New("parsing failed: more than 2 commas found")
	}
	x1, err := utils.ToInt(firstSplit[0])
	if err != nil {
		return line, err
	}
	line.Start.X = x1
	y2, err := utils.ToInt(firstSplit[2])
	if err != nil {
		return line, err
	}
	line.End.Y = y2
	secondSplit := strings.Split(firstSplit[1], " ")
	if len(secondSplit) != 3 {
		return line, errors.New("parsing failed: more than 2 spaces found")
	}
	y1, err := utils.ToInt(secondSplit[0])
	if err != nil {
		return line, err
	}
	line.Start.Y = y1
	x2, err := utils.ToInt(secondSplit[2])
	if err != nil {
		return line, err
	}
	line.End.X = x2
	return line, nil
}

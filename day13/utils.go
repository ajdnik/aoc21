package day13

import (
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

type Direction int

const (
	XAxis Direction = iota
	YAxis
)

type Fold struct {
	Axis      Direction
	Dimension int64
}

func ToFold(data string) (*Fold, error) {
	parts := strings.Split(data, "=")
	dir := XAxis
	if parts[0] == "fold along y" {
		dir = YAxis
	}
	dim, err := utils.ToInt(parts[1])
	if err != nil {
		return nil, err
	}
	return &Fold{
		Axis:      dir,
		Dimension: dim,
	}, nil
}

func ToPaperAndFolds(lines []string) ([][]bool, []*Fold, error) {
	dots := [][]int64{}
	var maxY, maxX int64
	folds := []*Fold{}
	for _, line := range lines {
		if strings.HasPrefix(line, "fold along") {
			fold, err := ToFold(line)
			if err != nil {
				return nil, nil, err
			}
			folds = append(folds, fold)
			continue
		}
		dot, err := utils.ToIntList(line, ",")
		if err != nil {
			return nil, nil, err
		}
		if dot[0] > maxX {
			maxX = dot[0]
		}
		if dot[1] > maxY {
			maxY = dot[1]
		}
		dots = append(dots, dot)
	}

	paper := make([][]bool, int(maxX+1))
	for row := 0; row < len(paper); row++ {
		paper[row] = make([]bool, int(maxY+1))
	}

	for _, dot := range dots {
		paper[dot[0]][dot[1]] = true
	}
	return paper, folds, nil
}

func FoldPaper(paper [][]bool, fold *Fold) [][]bool {
	if fold.Axis == YAxis {
		folded := make([][]bool, len(paper))
		for row := 0; row < len(folded); row++ {
			folded[row] = make([]bool, int(fold.Dimension))
			for col := 0; col < len(folded[row]); col++ {
				folded[row][col] = paper[row][col] || paper[row][len(paper[row])-1-col]
			}
		}
		return folded
	}
	folded := make([][]bool, int(fold.Dimension))
	for row := 0; row < len(folded); row++ {
		folded[row] = make([]bool, len(paper[row]))
		for col := 0; col < len(folded[row]); col++ {
			folded[row][col] = paper[row][col] || paper[len(paper)-1-row][col]
		}
	}
	return folded
}

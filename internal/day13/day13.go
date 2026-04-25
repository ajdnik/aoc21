package day13

import (
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

type direction int

const (
	xAxis direction = iota
	yAxis
)

type fold struct {
	axis      direction
	dimension int64
}

func parsePaperAndFolds(lines []string) ([][]bool, []*fold) {
	var dots [][]int64
	var maxY, maxX int64
	var folds []*fold

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "fold along") {
			parts := strings.Split(line, "=")
			dir := xAxis
			if parts[0] == "fold along y" {
				dir = yAxis
			}
			dim, err := utils.ToInt(parts[1])
			if err != nil {
				panic(err)
			}
			folds = append(folds, &fold{axis: dir, dimension: dim})
			continue
		}
		dot, err := utils.ToIntList(line, ",")
		if err != nil {
			panic(err)
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
	return paper, folds
}

func foldPaper(paper [][]bool, f *fold) [][]bool {
	if f.axis == yAxis {
		folded := make([][]bool, len(paper))
		for row := 0; row < len(folded); row++ {
			folded[row] = make([]bool, int(f.dimension))
			for col := 0; col < len(folded[row]); col++ {
				folded[row][col] = paper[row][col] || paper[row][len(paper[row])-1-col]
			}
		}
		return folded
	}
	folded := make([][]bool, int(f.dimension))
	for row := 0; row < len(folded); row++ {
		folded[row] = make([]bool, len(paper[row]))
		for col := 0; col < len(folded[row]); col++ {
			folded[row][col] = paper[row][col] || paper[len(paper)-1-row][col]
		}
	}
	return folded
}

func countDots(paper [][]bool) int64 {
	var count int64
	for row := 0; row < len(paper); row++ {
		for col := 0; col < len(paper[row]); col++ {
			if paper[row][col] {
				count++
			}
		}
	}
	return count
}

func Part1(lines []string) int64 {
	paper, folds := parsePaperAndFolds(lines)
	folded := foldPaper(paper, folds[0])
	return countDots(folded)
}

func Part2(lines []string) string {
	paper, folds := parsePaperAndFolds(lines)
	for _, f := range folds {
		paper = foldPaper(paper, f)
	}

	var sb strings.Builder
	for col := 0; col < len(paper[0]); col++ {
		if col > 0 {
			sb.WriteByte('\n')
		}
		for row := 0; row < len(paper); row++ {
			if paper[row][col] {
				sb.WriteByte('#')
			} else {
				sb.WriteByte('.')
			}
		}
	}
	return sb.String()
}

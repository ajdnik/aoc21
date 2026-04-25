// Package day13 solves AoC 2021 day 13: Transparent Origami.
// Fold a dot grid along axes and read the resulting letter pattern.
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
	dimension int
}

func parsePaperAndFolds(lines []string) ([][]bool, []*fold) {
	var dots [][2]int
	var maxY, maxX int
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
		coords, err := utils.ToIntList(line, ",")
		if err != nil {
			panic(err)
		}
		if coords[0] > maxX {
			maxX = coords[0]
		}
		if coords[1] > maxY {
			maxY = coords[1]
		}
		dots = append(dots, [2]int{coords[0], coords[1]})
	}

	paper := make([][]bool, maxX+1)
	for row := range paper {
		paper[row] = make([]bool, maxY+1)
	}
	for _, dot := range dots {
		paper[dot[0]][dot[1]] = true
	}
	return paper, folds
}

func foldPaper(paper [][]bool, f *fold) [][]bool {
	if f.axis == yAxis {
		folded := make([][]bool, len(paper))
		for row := range folded {
			folded[row] = make([]bool, f.dimension)
			for col := 0; col < f.dimension; col++ {
				mirror := 2*f.dimension - col
				folded[row][col] = paper[row][col]
				if mirror < len(paper[row]) {
					folded[row][col] = folded[row][col] || paper[row][mirror]
				}
			}
		}
		return folded
	}
	folded := make([][]bool, f.dimension)
	for row := 0; row < f.dimension; row++ {
		folded[row] = make([]bool, len(paper[row]))
		mirror := 2*f.dimension - row
		for col := 0; col < len(folded[row]); col++ {
			folded[row][col] = paper[row][col]
			if mirror < len(paper) {
				folded[row][col] = folded[row][col] || paper[mirror][col]
			}
		}
	}
	return folded
}

func countDots(paper [][]bool) int {
	var count int
	for _, row := range paper {
		for _, cell := range row {
			if cell {
				count++
			}
		}
	}
	return count
}

func Part1(lines []string) int {
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

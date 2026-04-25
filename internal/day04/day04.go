package day04

import (
	"slices"
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

type bingoBoard struct {
	Numbers  [5][5]int
	Selected [5][5]bool
}

func (b *bingoBoard) markNumber(num int) bool {
	for row := range b.Numbers {
		for col := range b.Numbers[row] {
			if b.Numbers[row][col] == num && !b.Selected[row][col] {
				b.Selected[row][col] = true
				return true
			}
		}
	}
	return false
}

func (b *bingoBoard) sumUnselected() int {
	var sum int
	for row := range b.Numbers {
		for col := range b.Numbers[row] {
			if !b.Selected[row][col] {
				sum += b.Numbers[row][col]
			}
		}
	}
	return sum
}

func (b *bingoBoard) hasBingo() bool {
	for i := range b.Selected {
		hasCol, hasRow := true, true
		for j := range b.Selected[i] {
			hasCol = hasCol && b.Selected[i][j]
			hasRow = hasRow && b.Selected[j][i]
		}
		if hasCol || hasRow {
			return true
		}
	}
	return false
}

func parseInput(lines []string) ([]int, []*bingoBoard) {
	sNums := strings.Split(lines[0], ",")
	drawNumbers := make([]int, len(sNums))
	for i, s := range sNums {
		num, err := utils.ToInt(s)
		if err != nil {
			panic(err)
		}
		drawNumbers[i] = num
	}

	var boards []*bingoBoard
	var buff [5]string
	var idx int
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		buff[idx] = line
		idx++
		if idx == 5 {
			board := &bingoBoard{}
			for row, l := range buff {
				l = utils.NormalizeSpaces(l)
				parts := strings.Split(l, " ")
				for col, part := range parts {
					num, err := utils.ToInt(part)
					if err != nil {
						panic(err)
					}
					board.Numbers[row][col] = num
				}
			}
			boards = append(boards, board)
			idx = 0
		}
	}
	return drawNumbers, boards
}

func Part1(lines []string) int {
	drawNumbers, boards := parseInput(lines)

	for _, draw := range drawNumbers {
		for _, board := range boards {
			board.markNumber(draw)
			if board.hasBingo() {
				return board.sumUnselected() * draw
			}
		}
	}
	return -1
}

func Part2(lines []string) int {
	drawNumbers, boards := parseInput(lines)

	var won []int
	var wonDraw []int
	for _, draw := range drawNumbers {
		for idx, board := range boards {
			if slices.Contains(won, idx) {
				continue
			}
			board.markNumber(draw)
			if board.hasBingo() {
				won = append(won, idx)
				wonDraw = append(wonDraw, draw)
			}
		}
	}

	last := len(won) - 1
	return boards[won[last]].sumUnselected() * wonDraw[last]
}

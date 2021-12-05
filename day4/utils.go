package day4

import (
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

type BingoBoard struct {
	Numbers  [5][5]int64
	Selected [5][5]bool
}

func (bb *BingoBoard) MarkNumber(num int64) bool {
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if bb.Numbers[row][col] == num && !bb.Selected[row][col] {
				bb.Selected[row][col] = true
				return true
			}
		}
	}
	return false
}

func (bb *BingoBoard) SumUnselected() int64 {
	sum := int64(0)
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !bb.Selected[row][col] {
				sum += bb.Numbers[row][col]
			}
		}
	}
	return sum
}

func (bb *BingoBoard) HasBingo() bool {
	for i := 0; i < 5; i++ {
		hasCol, hasRow := true, true
		for j := 0; j < 5; j++ {
			hasCol = hasCol && bb.Selected[i][j]
			hasRow = hasRow && bb.Selected[j][i]
		}
		if hasCol || hasRow {
			return true
		}
	}
	return false
}

func ToDrawNumbers(data string) ([]int64, error) {
	sNums := strings.Split(data, ",")
	iNums := []int64{}
	for _, itm := range sNums {
		num, err := utils.ToInt(itm)
		if err != nil {
			return nil, err
		}
		iNums = append(iNums, num)
	}
	return iNums, nil
}

func ToBingoBoard(data [5]string) (*BingoBoard, error) {
	board := &BingoBoard{}
	for row, line := range data {
		line = utils.NormalizeSpaces(line)
		sNums := strings.Split(line, " ")
		for col, itm := range sNums {
			num, err := utils.ToInt(itm)
			if err != nil {
				return nil, err
			}
			board.Numbers[row][col] = num
		}
	}
	return board, nil
}

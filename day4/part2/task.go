package main

import (
	"log"

	"github.com/ajdnik/aoc21/day4"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	var drawNumbers []int64
	var boards []*day4.BingoBoard
	var idx int
	first := true
	buff := [5]string{}
	for scanner.Scan() {
		data := scanner.Text()

		if data == "" {
			continue
		}

		if first {
			drawNumbers, err = day4.ToDrawNumbers(data)
			if err != nil {
				log.Fatal(err)
			}
			first = false
			continue
		}

		buff[idx] = data
		idx++

		if idx == 5 {
			board, err := day4.ToBingoBoard(buff)
			if err != nil {
				log.Fatal(err)
			}
			boards = append(boards, board)
			idx = 0
		}
	}

	won := []int{}
	wonDraw := []int64{}
	for _, draw := range drawNumbers {
		for idx, board := range boards {
			if utils.IsIncluded(won, idx) {
				continue
			}
			board.MarkNumber(draw)
			if board.HasBingo() {
				won = append(won, idx)
				wonDraw = append(wonDraw, draw)
			}
		}
	}

	last := len(won) - 1
	sum := boards[won[last]].SumUnselected()
	log.Printf("bingo board=%d, sum=%d, draw=%d, score=%d\n", won[last], sum, wonDraw[last], sum*wonDraw[last])
}

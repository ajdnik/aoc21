package main

import (
	"log"

	"github.com/ajdnik/aoc21/day13"
	"github.com/ajdnik/aoc21/utils"
)

func CountDots(paper [][]bool) int64 {
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

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	paper, folds, err := day13.ToPaperAndFolds(lines)
	if err != nil {
		log.Fatal(err)
	}

	folded := day13.FoldPaper(paper, folds[0])
	dots := CountDots(folded)
	log.Printf("dots=%d\n", dots)
}

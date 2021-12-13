package main

import (
	"log"

	"github.com/ajdnik/aoc21/day13"
	"github.com/ajdnik/aoc21/utils"
)

func PrintPaper(paper [][]bool) {
	for col := 0; col < len(paper[0]); col++ {
		line := ""
		for row := 0; row < len(paper); row++ {
			if paper[row][col] {
				line += "#"
			} else {
				line += "."
			}
		}
		log.Printf("%s\n", line)
	}
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

	for _, fold := range folds {
		paper = day13.FoldPaper(paper, fold)
	}

	PrintPaper(paper)
}

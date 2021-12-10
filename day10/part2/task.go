package main

import (
	"log"
	"sort"

	"github.com/ajdnik/aoc21/day10"
	"github.com/ajdnik/aoc21/utils"
)

const IncorrectLine = -1

var Lut = map[rune]int64{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func GetLineScore(line string) int64 {
	stack, char := day10.ParseLine(line)
	if char != day10.EmptyRune {
		return IncorrectLine
	}
	var score int64
	var ok bool
	for true {
		stack, char, ok = day10.Pop(stack)
		if !ok {
			break
		}
		score *= 5
		score += Lut[char]
	}
	return score

}

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	scores := []int64{}
	for scanner.Scan() {
		data := scanner.Text()
		score := GetLineScore(data)
		if score == IncorrectLine {
			continue
		}
		scores = append(scores, score)
	}

	sort.Slice(scores, func(i, j int) bool { return scores[i] > scores[j] })
	log.Printf("middle=%d\n", scores[len(scores)/2])
}

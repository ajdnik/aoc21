package main

import (
	"log"

	"github.com/ajdnik/aoc21/day10"
	"github.com/ajdnik/aoc21/utils"
)

var Lut = map[rune]int64{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func GetLineScore(line string) int64 {
	_, char := day10.ParseLine(line)
	if val, ok := Lut[char]; ok {
		return val
	}
	return 0
}

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	var score int64
	for scanner.Scan() {
		data := scanner.Text()
		score += GetLineScore(data)
	}

	log.Printf("score=%d\n", score)
}

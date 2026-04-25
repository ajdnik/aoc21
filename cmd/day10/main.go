package main

import (
	"fmt"
	"log"

	"github.com/ajdnik/aoc21/internal/day10"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	lines, err := utils.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", day10.Part1(lines))
	fmt.Println("Part 2:", day10.Part2(lines))
}

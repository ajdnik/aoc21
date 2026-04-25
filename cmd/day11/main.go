package main

import (
	"fmt"
	"log"

	"github.com/ajdnik/aoc21/internal/day11"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	lines, err := utils.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", day11.Part1(lines))
	fmt.Println("Part 2:", day11.Part2(lines))
}

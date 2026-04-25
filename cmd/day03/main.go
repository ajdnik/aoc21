package main

import (
	"fmt"
	"log"

	"github.com/ajdnik/aoc21/internal/day03"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	lines, err := utils.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", day03.Part1(lines))
	fmt.Println("Part 2:", day03.Part2(lines))
}

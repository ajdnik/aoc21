package main

import (
	"fmt"
	"log"

	"github.com/ajdnik/aoc21/internal/day13"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	lines, err := utils.ReadLines()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", day13.Part1(lines))
	fmt.Println("Part 2:")
	fmt.Println(day13.Part2(lines))
}

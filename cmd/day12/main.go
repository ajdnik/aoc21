package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ajdnik/aoc21/internal/day12"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: day12 <input-file>")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() //nolint:errcheck

	lines := utils.ReadLines(f)
	fmt.Println("Part 1:", day12.Part1(lines))
	fmt.Println("Part 2:", day12.Part2(lines))
}

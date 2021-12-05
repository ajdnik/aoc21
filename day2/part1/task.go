package main

import (
	"log"

	"github.com/ajdnik/aoc21/day2"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	var horiz, depth int64
	for scanner.Scan() {
		mov, err := day2.ToMovement(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		switch mov.Direction {
		case day2.Forward:
			horiz += mov.Unit
		case day2.Up:
			depth -= mov.Unit
		case day2.Down:
			depth += mov.Unit
		default:
			log.Fatal("unknown direction")
		}
	}

	log.Printf("horizontal=%d, depth=%d, product=%d.\n", horiz, depth, horiz*depth)
}

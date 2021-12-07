package main

import (
	"log"

	"github.com/ajdnik/aoc21/day6"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	var timers []int64
	for scanner.Scan() {
		data := scanner.Text()
		timers, err = utils.ToIntList(data, ",")
		if err != nil {
			log.Fatal(err)
		}
	}

	timers = day6.SimulateDays(timers, 80)

	log.Printf("fish=%d\n", len(timers))
}

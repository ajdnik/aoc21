package main

import (
	"log"

	"github.com/ajdnik/aoc21/day11"
	"github.com/ajdnik/aoc21/utils"
)

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	energies := [][]int64{}
	for scanner.Scan() {
		data := scanner.Text()
		levels, err := utils.ToIntList(data, "")
		if err != nil {
			log.Fatal(err)
		}
		energies = append(energies, levels)
	}

	var flashes, total int64
	for step := 0; step < 100; step++ {
		energies, flashes = day11.SimulateStep(energies)
		total += flashes
	}
	log.Printf("flashes=%d\n", total)
}

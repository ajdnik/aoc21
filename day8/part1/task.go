package main

import (
	"log"

	"github.com/ajdnik/aoc21/day8"
	"github.com/ajdnik/aoc21/utils"
)

func CountDigits(observations []*day8.Observation) uint64 {
	var sum uint64
	for _, observ := range observations {
		for _, out := range observ.Output {
			switch len(out) {
			case 2:
				sum++
			case 3:
				sum++
			case 4:
				sum++
			case 7:
				sum++
			}
		}
	}
	return sum
}

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	observations := []*day8.Observation{}
	for scanner.Scan() {
		data := scanner.Text()
		observation := day8.ToObservation(data)
		observations = append(observations, observation)
	}

	digitCounts := CountDigits(observations)
	log.Printf("digitCount=%d\n", digitCounts)
}

package main

import (
	"log"

	"github.com/ajdnik/aoc21/utils"
)

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	var ones, zeros []uint64
	for scanner.Scan() {
		data := scanner.Text()
		if ones == nil {
			ones = make([]uint64, len(data))
			zeros = make([]uint64, len(data))
		}

		for idx, char := range data {
			switch char {
			case '1':
				ones[idx]++
			case '0':
				zeros[idx]++
			default:
				log.Fatal("unknown binary character")
			}
		}
	}

	var gamma, epsilon string
	for i := 0; i < len(zeros); i++ {
		if zeros[i] > ones[i] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaInt, err := utils.BinaryToInt(gamma)
	if err != nil {
		log.Fatal(err)
	}

	epsilonInt, err := utils.BinaryToInt(epsilon)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("gamma=%s, gammaInt=%d, epsilon=%s, epsilonInt=%d, power=%d\n", gamma, gammaInt, epsilon, epsilonInt, gammaInt*epsilonInt)
}

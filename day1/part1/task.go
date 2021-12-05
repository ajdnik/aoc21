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

	first := true
	var prev int64
	var inc uint64
	for scanner.Scan() {
		num, err := utils.ToInt(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if first {
			prev = num
			first = false
			continue
		}

		if prev < num {
			inc++
		}
		prev = num
	}

	log.Printf("increased=%d\n", inc)
}

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

	var first, inc uint
	var prev int64
	window := make([]int64, 3)
	for scanner.Scan() {
		num, err := utils.ToInt(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if first < 3 {
			window[first] = num
			first++
			if first == 3 {
				prev = utils.Sum(window)
				window[0] = window[1]
				window[1] = window[2]
			}
			continue
		}

		window[2] = num
		sum := utils.Sum(window)
		window[0] = window[1]
		window[1] = window[2]

		if prev < sum {
			inc++
		}
		prev = sum
	}

	log.Printf("increased=%d\n", inc)
}

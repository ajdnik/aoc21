package main

import (
	"log"

	"github.com/ajdnik/aoc21/day6"
	"github.com/ajdnik/aoc21/utils"
)

func logProgress(iter int, total int) {
	if iter%25 != 0 {
		return
	}
	log.Printf("processed=%d, total=%d\n", iter, total)
}

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	var initialTimers []int64
	for scanner.Scan() {
		data := scanner.Text()
		initialTimers, err = utils.ToIntList(data, ",")
		if err != nil {
			log.Fatal(err)
		}
	}

	initLen := len(initialTimers)
	cache := map[int64]uint64{}
	var totalFishes uint64
	for iter, initialTimer := range initialTimers {
		if val, ok := cache[initialTimer]; ok {
			totalFishes += val
			logProgress(iter, initLen)
			continue
		}

		timers := []int64{initialTimer}
		timers = day6.SimulateDays(timers, 256)
		popLen := uint64(len(timers))
		cache[initialTimer] = popLen
		totalFishes += popLen
		logProgress(iter, initLen)
	}

	log.Printf("fish=%d\n", totalFishes)
}

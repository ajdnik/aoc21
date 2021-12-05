package main

import (
	"errors"
	"log"

	"github.com/ajdnik/aoc21/utils"
)

type SearchType int

const (
	Oxygen SearchType = iota
	CO2
)

func FilterMatches(data []string, match byte, pos int) []string {
	filtered := []string{}
	for _, itm := range data {
		if itm[pos] == match {
			filtered = append(filtered, itm)
		}
	}
	return filtered
}

func BuildValue(data []string, typ SearchType, pos int) (string, error) {
	var zeros, ones uint64
	for _, itm := range data {
		switch itm[pos] {
		case '1':
			ones++
		case '0':
			zeros++
		default:
			return "", errors.New("unknown binary character")
		}
	}

	var filtered []string
	switch typ {
	case Oxygen:
		if ones >= zeros {
			filtered = FilterMatches(data, '1', pos)
		} else {
			filtered = FilterMatches(data, '0', pos)
		}
	case CO2:
		if ones < zeros {
			filtered = FilterMatches(data, '1', pos)
		} else {
			filtered = FilterMatches(data, '0', pos)
		}
	default:
		return "", errors.New("unsupported search type")
	}

	if len(filtered) == 1 {
		return filtered[0], nil
	}
	pos++
	return BuildValue(filtered, typ, pos)
}

func main() {
	scanner, closer, err := utils.ScanFile()
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	data := []string{}
	for scanner.Scan() {
		itm := scanner.Text()
		data = append(data, itm)
	}

	oxygen, err := BuildValue(data, Oxygen, 0)
	if err != nil {
		log.Fatal(err)
	}

	co2, err := BuildValue(data, CO2, 0)
	if err != nil {
		log.Fatal(err)
	}

	oxygenInt, err := utils.BinaryToInt(oxygen)
	if err != nil {
		log.Fatal(err)
	}

	co2Int, err := utils.BinaryToInt(co2)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("oxygen=%s, oxygenInt=%d, co2=%s, co2Int=%d, life=%d\n", oxygen, oxygenInt, co2, co2Int, oxygenInt*co2Int)
}

package day8

import (
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

type Observation struct {
	Patterns []string
	Output   []string
}

func ToObservation(data string) *Observation {
	data = utils.NormalizeSpaces(data)
	split := strings.Split(data, "|")
	patterns := strings.Split(split[0], " ")
	output := strings.Split(split[1], " ")
	return &Observation{
		Patterns: patterns,
		Output:   output,
	}
}

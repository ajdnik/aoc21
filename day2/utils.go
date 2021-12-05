package day2

import (
	"errors"
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

type Direction int

const (
	Forward Direction = iota
	Up
	Down
)

type Movement struct {
	Direction Direction
	Unit      int64
}

func ToMovement(data string) (Movement, error) {
	var mov Movement
	data = utils.NormalizeSpaces(data)
	parts := strings.Split(data, " ")
	if len(parts) != 2 {
		return mov, errors.New("movement input needs two elements")
	}

	switch strings.ToLower(parts[0]) {
	case "forward":
		mov.Direction = Forward
	case "up":
		mov.Direction = Up
	case "down":
		mov.Direction = Down
	default:
		return mov, errors.New("unknown movement direction")
	}

	unit, err := utils.ToInt(parts[1])
	if err != nil {
		return mov, err
	}
	mov.Unit = unit
	return mov, nil
}

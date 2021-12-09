package day9

import (
	"strings"

	"github.com/ajdnik/aoc21/utils"
)

func ToHeights(data string) ([]int64, error) {
	parts := strings.Split(data, "")
	heights := []int64{}
	for _, part := range parts {
		height, err := utils.ToInt(part)
		if err != nil {
			return nil, err
		}
		heights = append(heights, height)
	}
	return heights, nil
}

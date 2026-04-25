package day11

import (
	"github.com/ajdnik/aoc21/utils"
)

func parseEnergies(lines []string) [][]int64 {
	energies := make([][]int64, len(lines))
	for i, line := range lines {
		levels, err := utils.ToIntList(line, "")
		if err != nil {
			panic(err)
		}
		energies[i] = levels
	}
	return energies
}

func checkFlash(energies [][]int64, flashes [][]bool, row, col int, inc int64) {
	if row < 0 || row >= len(energies) {
		return
	}
	if col < 0 || col >= len(energies[row]) {
		return
	}
	energies[row][col] += inc
	if energies[row][col] <= 9 || flashes[row][col] {
		return
	}
	flashes[row][col] = true
	checkFlash(energies, flashes, row+1, col, 1)
	checkFlash(energies, flashes, row-1, col, 1)
	checkFlash(energies, flashes, row, col+1, 1)
	checkFlash(energies, flashes, row, col-1, 1)
	checkFlash(energies, flashes, row-1, col-1, 1)
	checkFlash(energies, flashes, row+1, col+1, 1)
	checkFlash(energies, flashes, row+1, col-1, 1)
	checkFlash(energies, flashes, row-1, col+1, 1)
}

func simulateStep(energies [][]int64) int64 {
	flashes := make([][]bool, len(energies))
	for row := 0; row < len(energies); row++ {
		energies[row] = utils.Add(energies[row], 1)
		flashes[row] = make([]bool, len(energies[row]))
	}

	for row := 0; row < len(energies); row++ {
		for col := 0; col < len(energies[row]); col++ {
			checkFlash(energies, flashes, row, col, 0)
		}
	}

	var total int64
	for row := 0; row < len(energies); row++ {
		for col := 0; col < len(energies[row]); col++ {
			if flashes[row][col] {
				total++
				energies[row][col] = 0
			}
		}
	}
	return total
}

func Part1(lines []string) int64 {
	energies := parseEnergies(lines)
	var total int64
	for step := 0; step < 100; step++ {
		total += simulateStep(energies)
	}
	return total
}

func Part2(lines []string) int64 {
	energies := parseEnergies(lines)
	for step := int64(1); ; step++ {
		if simulateStep(energies) == 100 {
			return step
		}
	}
}

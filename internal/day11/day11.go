package day11

import (
	"github.com/ajdnik/aoc21/utils"
)

func parseEnergies(lines []string) [][]int {
	return utils.ParseDigitGrid(lines)
}

func checkFlash(energies [][]int, flashes [][]bool, row, col int, inc int) {
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

func simulateStep(energies [][]int) int {
	flashes := make([][]bool, len(energies))
	for row := 0; row < len(energies); row++ {
		for i := range energies[row] {
			energies[row][i]++
		}
		flashes[row] = make([]bool, len(energies[row]))
	}

	for row := 0; row < len(energies); row++ {
		for col := 0; col < len(energies[row]); col++ {
			checkFlash(energies, flashes, row, col, 0)
		}
	}

	var total int
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

func Part1(lines []string) int {
	energies := parseEnergies(lines)
	var total int
	for step := 0; step < 100; step++ {
		total += simulateStep(energies)
	}
	return total
}

func Part2(lines []string) int {
	energies := parseEnergies(lines)
	for step := 1; ; step++ {
		if simulateStep(energies) == 100 {
			return step
		}
	}
}

package day11

import "github.com/ajdnik/aoc21/utils"

func CheckFlash(energies [][]int64, flashes [][]bool, row, col int, inc int64) ([][]int64, [][]bool) {
	if row < 0 || row >= len(energies) {
		return energies, flashes
	}
	if col < 0 || col >= len(energies[row]) {
		return energies, flashes
	}
	energies[row][col] += inc
	if energies[row][col] <= 9 {
		return energies, flashes
	}
	if flashes[row][col] {
		return energies, flashes
	}
	flashes[row][col] = true
	energies, flashes = CheckFlash(energies, flashes, row+1, col, 1)
	energies, flashes = CheckFlash(energies, flashes, row-1, col, 1)
	energies, flashes = CheckFlash(energies, flashes, row, col+1, 1)
	energies, flashes = CheckFlash(energies, flashes, row, col-1, 1)
	energies, flashes = CheckFlash(energies, flashes, row-1, col-1, 1)
	energies, flashes = CheckFlash(energies, flashes, row+1, col+1, 1)
	energies, flashes = CheckFlash(energies, flashes, row+1, col-1, 1)
	energies, flashes = CheckFlash(energies, flashes, row-1, col+1, 1)
	return energies, flashes
}

func SimulateStep(energies [][]int64) ([][]int64, int64) {
	flashes := make([][]bool, len(energies))
	for row := 0; row < len(energies); row++ {
		energies[row] = utils.Add(energies[row], 1)
		flashes[row] = make([]bool, len(energies[row]))
	}

	for row := 0; row < len(energies); row++ {
		for col := 0; col < len(energies[row]); col++ {
			energies, flashes = CheckFlash(energies, flashes, row, col, 0)
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

	return energies, total
}

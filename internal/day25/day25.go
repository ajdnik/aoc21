package day25

func Part1(lines []string) int {
	rows := len(lines)
	cols := len(lines[0])
	grid := make([][]byte, rows)
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	for step := 1; ; step++ {
		moved := false

		// East-facing move
		next := make([][]byte, rows)
		for r := range grid {
			next[r] = make([]byte, cols)
			copy(next[r], grid[r])
		}
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] == '>' {
					nc := (c + 1) % cols
					if grid[r][nc] == '.' {
						next[r][nc] = '>'
						next[r][c] = '.'
						moved = true
					}
				}
			}
		}
		grid = next

		// South-facing move
		next2 := make([][]byte, rows)
		for r := range grid {
			next2[r] = make([]byte, cols)
			copy(next2[r], grid[r])
		}
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] == 'v' {
					nr := (r + 1) % rows
					if grid[nr][c] == '.' {
						next2[nr][c] = 'v'
						next2[r][c] = '.'
						moved = true
					}
				}
			}
		}
		grid = next2

		if !moved {
			return step
		}
	}
}

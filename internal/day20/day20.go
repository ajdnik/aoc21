// Package day20 solves AoC 2021 day 20: Trench Map.
// Apply an image enhancement algorithm, handling the infinite background flip.
package day20

type image struct {
	pixels     map[[2]int]bool
	minR, maxR int
	minC, maxC int
	bgLit      bool
}

func parseInput(lines []string) (string, *image) {
	algo := lines[0]
	img := &image{pixels: map[[2]int]bool{}}
	for r, line := range lines[2:] {
		for c, ch := range line {
			if ch == '#' {
				img.pixels[[2]int{r, c}] = true
			}
			if c > img.maxC {
				img.maxC = c
			}
		}
		if r > img.maxR {
			img.maxR = r
		}
	}
	return algo, img
}

// enhance applies one step: for each pixel, read 3x3 neighborhood as 9-bit index into algo.
func enhance(algo string, img *image) *image {
	next := &image{
		pixels: map[[2]int]bool{},
		minR:   img.minR - 1,
		maxR:   img.maxR + 1,
		minC:   img.minC - 1,
		maxC:   img.maxC + 1,
	}

	for r := next.minR; r <= next.maxR; r++ {
		for c := next.minC; c <= next.maxC; c++ {
			idx := 0
			for dr := -1; dr <= 1; dr++ {
				for dc := -1; dc <= 1; dc++ {
					idx <<= 1
					nr, nc := r+dr, c+dc
					lit := false
					if nr >= img.minR && nr <= img.maxR && nc >= img.minC && nc <= img.maxC {
						lit = img.pixels[[2]int{nr, nc}]
					} else {
						lit = img.bgLit
					}
					if lit {
						idx |= 1
					}
				}
			}
			if algo[idx] == '#' {
				next.pixels[[2]int{r, c}] = true
			}
		}
	}

	if img.bgLit {
		next.bgLit = algo[511] == '#'
	} else {
		next.bgLit = algo[0] == '#'
	}

	return next
}

func solve(lines []string, steps int) int {
	algo, img := parseInput(lines)
	for range steps {
		img = enhance(algo, img)
	}
	return len(img.pixels)
}

func Part1(lines []string) int {
	return solve(lines, 2)
}

func Part2(lines []string) int {
	return solve(lines, 50)
}

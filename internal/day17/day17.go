package day17

import (
	"fmt"
)

type target struct {
	x1, x2, y1, y2 int
}

func parseTarget(line string) target {
	var t target
	fmt.Sscanf(line, "target area: x=%d..%d, y=%d..%d", &t.x1, &t.x2, &t.y1, &t.y2)
	return t
}

func simulate(vx, vy int, t target) (int, bool) {
	x, y, maxY := 0, 0, 0
	for {
		x += vx
		y += vy
		if y > maxY {
			maxY = y
		}
		if vx > 0 {
			vx--
		} else if vx < 0 {
			vx++
		}
		vy--
		if x >= t.x1 && x <= t.x2 && y >= t.y1 && y <= t.y2 {
			return maxY, true
		}
		if x > t.x2 || y < t.y1 {
			return 0, false
		}
	}
}

func Part1(lines []string) int {
	t := parseTarget(lines[0])
	best := 0
	for vx := 0; vx <= t.x2; vx++ {
		for vy := t.y1; vy <= -t.y1; vy++ {
			if maxY, hit := simulate(vx, vy, t); hit && maxY > best {
				best = maxY
			}
		}
	}
	return best
}

func Part2(lines []string) int {
	t := parseTarget(lines[0])
	count := 0
	for vx := 0; vx <= t.x2; vx++ {
		for vy := t.y1; vy <= -t.y1; vy++ {
			if _, hit := simulate(vx, vy, t); hit {
				count++
			}
		}
	}
	return count
}

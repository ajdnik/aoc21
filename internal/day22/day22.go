package day22

import (
	"fmt"
	"strings"
)

type cuboid struct {
	on                     bool
	x1, x2, y1, y2, z1, z2 int
}

func parseCuboids(lines []string) []cuboid {
	var cuboids []cuboid
	for _, line := range lines {
		var c cuboid
		if strings.HasPrefix(line, "on") {
			c.on = true
		}
		idx := strings.Index(line, "x=")
		if idx < 0 {
			panic("invalid input: " + line)
		}
		_, err := fmt.Sscanf(line[idx:], "x=%d..%d,y=%d..%d,z=%d..%d",
			&c.x1, &c.x2, &c.y1, &c.y2, &c.z1, &c.z2)
		if err != nil {
			panic(err)
		}
		cuboids = append(cuboids, c)
	}
	return cuboids
}

func intersect(a, b cuboid) (cuboid, bool) {
	c := cuboid{
		x1: max(a.x1, b.x1), x2: min(a.x2, b.x2),
		y1: max(a.y1, b.y1), y2: min(a.y2, b.y2),
		z1: max(a.z1, b.z1), z2: min(a.z2, b.z2),
	}
	if c.x1 > c.x2 || c.y1 > c.y2 || c.z1 > c.z2 {
		return c, false
	}
	return c, true
}

func volume(c cuboid) int {
	return (c.x2 - c.x1 + 1) * (c.y2 - c.y1 + 1) * (c.z2 - c.z1 + 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Inclusion-exclusion approach: maintain list of signed cuboids
type signedCuboid struct {
	cuboid
	sign int // +1 or -1
}

func solve(cuboids []cuboid) int {
	var active []signedCuboid

	for _, c := range cuboids {
		var toAdd []signedCuboid
		for _, a := range active {
			if inter, ok := intersect(c, a.cuboid); ok {
				// Cancel out overlap
				toAdd = append(toAdd, signedCuboid{inter, -a.sign})
			}
		}
		if c.on {
			toAdd = append(toAdd, signedCuboid{c, 1})
		}
		active = append(active, toAdd...)
	}

	total := 0
	for _, a := range active {
		total += a.sign * volume(a.cuboid)
	}
	return total
}

func Part1(lines []string) int {
	cuboids := parseCuboids(lines)
	bound := cuboid{x1: -50, x2: 50, y1: -50, y2: 50, z1: -50, z2: 50}
	var filtered []cuboid
	for _, c := range cuboids {
		if inter, ok := intersect(c, bound); ok {
			inter.on = c.on
			filtered = append(filtered, inter)
		}
	}
	return solve(filtered)
}

func Part2(lines []string) int {
	return solve(parseCuboids(lines))
}

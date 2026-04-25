package day19

import (
	"fmt"
	"strings"
)

type vec3 struct {
	x, y, z int
}

func (a vec3) sub(b vec3) vec3 { return vec3{a.x - b.x, a.y - b.y, a.z - b.z} }
func (a vec3) add(b vec3) vec3 { return vec3{a.x + b.x, a.y + b.y, a.z + b.z} }

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func manhattan(a, b vec3) int {
	return abs(a.x-b.x) + abs(a.y-b.y) + abs(a.z-b.z)
}

// All 24 rotation functions
var rotations = [24]func(vec3) vec3{
	func(v vec3) vec3 { return vec3{v.x, v.y, v.z} },
	func(v vec3) vec3 { return vec3{v.x, -v.z, v.y} },
	func(v vec3) vec3 { return vec3{v.x, -v.y, -v.z} },
	func(v vec3) vec3 { return vec3{v.x, v.z, -v.y} },
	func(v vec3) vec3 { return vec3{-v.x, -v.y, v.z} },
	func(v vec3) vec3 { return vec3{-v.x, v.z, v.y} },
	func(v vec3) vec3 { return vec3{-v.x, v.y, -v.z} },
	func(v vec3) vec3 { return vec3{-v.x, -v.z, -v.y} },
	func(v vec3) vec3 { return vec3{v.y, v.z, v.x} },
	func(v vec3) vec3 { return vec3{v.y, -v.x, v.z} },
	func(v vec3) vec3 { return vec3{v.y, -v.z, -v.x} },
	func(v vec3) vec3 { return vec3{v.y, v.x, -v.z} },
	func(v vec3) vec3 { return vec3{-v.y, -v.z, v.x} },
	func(v vec3) vec3 { return vec3{-v.y, v.x, v.z} },
	func(v vec3) vec3 { return vec3{-v.y, v.z, -v.x} },
	func(v vec3) vec3 { return vec3{-v.y, -v.x, -v.z} },
	func(v vec3) vec3 { return vec3{v.z, v.x, v.y} },
	func(v vec3) vec3 { return vec3{v.z, -v.y, v.x} },
	func(v vec3) vec3 { return vec3{v.z, -v.x, -v.y} },
	func(v vec3) vec3 { return vec3{v.z, v.y, -v.x} },
	func(v vec3) vec3 { return vec3{-v.z, -v.x, v.y} },
	func(v vec3) vec3 { return vec3{-v.z, v.y, v.x} },
	func(v vec3) vec3 { return vec3{-v.z, v.x, -v.y} },
	func(v vec3) vec3 { return vec3{-v.z, -v.y, -v.x} },
}

func parseScanners(lines []string) [][]vec3 {
	var scanners [][]vec3
	var current []vec3
	for _, line := range lines {
		if strings.HasPrefix(line, "---") {
			if current != nil {
				scanners = append(scanners, current)
			}
			current = []vec3{}
			continue
		}
		if line == "" {
			continue
		}
		var v vec3
		_, err := fmt.Sscanf(line, "%d,%d,%d", &v.x, &v.y, &v.z)
		if err != nil {
			panic(err)
		}
		current = append(current, v)
	}
	if current != nil {
		scanners = append(scanners, current)
	}
	return scanners
}

// tryMatch attempts to match scanner b against known beacons.
// Returns offset and rotated beacons if 12+ overlap found.
func tryMatch(known map[vec3]bool, beacons []vec3) (vec3, []vec3, bool) {
	knownSlice := make([]vec3, 0, len(known))
	for k := range known {
		knownSlice = append(knownSlice, k)
	}

	for _, rot := range rotations {
		rotated := make([]vec3, len(beacons))
		for i, b := range beacons {
			rotated[i] = rot(b)
		}
		// Try each pair of known beacon and rotated beacon as anchor
		for _, kb := range knownSlice {
			for _, rb := range rotated {
				offset := kb.sub(rb)
				matches := 0
				for _, r := range rotated {
					if known[r.add(offset)] {
						matches++
						if matches >= 12 {
							result := make([]vec3, len(rotated))
							for i, r2 := range rotated {
								result[i] = r2.add(offset)
							}
							return offset, result, true
						}
					}
				}
			}
		}
	}
	return vec3{}, nil, false
}

func solve(lines []string) (int, int) {
	scanners := parseScanners(lines)

	known := map[vec3]bool{}
	for _, b := range scanners[0] {
		known[b] = true
	}

	matched := make([]bool, len(scanners))
	matched[0] = true
	positions := []vec3{{0, 0, 0}}
	remaining := len(scanners) - 1

	for remaining > 0 {
		for i := 1; i < len(scanners); i++ {
			if matched[i] {
				continue
			}
			offset, beacons, ok := tryMatch(known, scanners[i])
			if !ok {
				continue
			}
			matched[i] = true
			remaining--
			positions = append(positions, offset)
			for _, b := range beacons {
				known[b] = true
			}
		}
	}

	maxDist := 0
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			if d := manhattan(positions[i], positions[j]); d > maxDist {
				maxDist = d
			}
		}
	}

	return len(known), maxDist
}

func Part1(lines []string) int {
	beacons, _ := solve(lines)
	return beacons
}

func Part2(lines []string) int {
	_, dist := solve(lines)
	return dist
}

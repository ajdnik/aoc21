package day24

import (
	"strconv"
	"strings"
)

// Each of the 14 blocks has this structure:
//   inp w
//   mul x 0
//   add x z
//   mod x 26
//   div z [divZ]    <- 1 or 26
//   add x [addX]    <- varies
//   eql x w
//   eql x 0
//   mul y 0
//   add y 25
//   mul y x
//   add y 1
//   mul z y
//   mul y 0
//   add y w
//   add y [addY]    <- varies
//   mul y x
//   add z y
//
// When divZ=1: push (digit + addY) onto base-26 stack
// When divZ=26: pop from stack, must match digit + addX = popped value

type blockParams struct {
	divZ, addX, addY int
}

func extractParams(lines []string) []blockParams {
	var params []blockParams
	for i := 0; i < len(lines); i += 18 {
		divZ := mustAtoi(strings.Fields(lines[i+4])[2])
		addX := mustAtoi(strings.Fields(lines[i+5])[2])
		addY := mustAtoi(strings.Fields(lines[i+15])[2])
		params = append(params, blockParams{divZ, addX, addY})
	}
	return params
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

// Solve finds digit constraints by pairing push/pop blocks.
// For each push(i)/pop(j) pair: digit[j] + addX[j] == digit[i] + addY[i]
// i.e. digit[j] = digit[i] + addY[i] + addX[j]
// So diff = addY[i] + addX[j], and digit[j] = digit[i] + diff
func solveConstraints(params []blockParams, maximize bool) [14]int {
	digits := [14]int{}
	var stack [][2]int // [index, addY]

	for j, p := range params {
		if p.divZ == 1 {
			// Push
			stack = append(stack, [2]int{j, p.addY})
		} else {
			// Pop
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i := top[0]
			diff := top[1] + p.addX // digit[j] = digit[i] + diff

			if maximize {
				if diff > 0 {
					digits[i] = 9 - diff
					digits[j] = 9
				} else {
					digits[i] = 9
					digits[j] = 9 + diff
				}
			} else {
				if diff > 0 {
					digits[i] = 1
					digits[j] = 1 + diff
				} else {
					digits[i] = 1 - diff
					digits[j] = 1
				}
			}
		}
	}
	return digits
}

func digitsToInt(digits [14]int) int {
	result := 0
	for _, d := range digits {
		result = result*10 + d
	}
	return result
}

// verify runs the MONAD program with the given digits and checks z==0
func verify(lines []string, digits [14]int) bool {
	vars := map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}
	inputIdx := 0

	val := func(s string) int {
		if s == "w" || s == "x" || s == "y" || s == "z" {
			return vars[s]
		}
		return mustAtoi(s)
	}

	for _, line := range lines {
		parts := strings.Fields(line)
		switch parts[0] {
		case "inp":
			vars[parts[1]] = digits[inputIdx]
			inputIdx++
		case "add":
			vars[parts[1]] += val(parts[2])
		case "mul":
			vars[parts[1]] *= val(parts[2])
		case "div":
			vars[parts[1]] /= val(parts[2])
		case "mod":
			vars[parts[1]] %= val(parts[2])
		case "eql":
			if vars[parts[1]] == val(parts[2]) {
				vars[parts[1]] = 1
			} else {
				vars[parts[1]] = 0
			}
		}
	}
	return vars["z"] == 0
}

func Part1(lines []string) int {
	params := extractParams(lines)
	digits := solveConstraints(params, true)
	if !verify(lines, digits) {
		panic("verification failed for part 1")
	}
	return digitsToInt(digits)
}

func Part2(lines []string) int {
	params := extractParams(lines)
	digits := solveConstraints(params, false)
	if !verify(lines, digits) {
		panic("verification failed for part 2")
	}
	return digitsToInt(digits)
}

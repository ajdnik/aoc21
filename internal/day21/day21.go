// Package day21 solves AoC 2021 day 21: Dirac Dice.
// Simulate deterministic and quantum dice games on a circular board.
package day21

import (
	"fmt"
)

func parseStarts(lines []string) (int, int) {
	var p1, p2 int
	if _, err := fmt.Sscanf(lines[0], "Player 1 starting position: %d", &p1); err != nil {
		panic(err)
	}
	if _, err := fmt.Sscanf(lines[1], "Player 2 starting position: %d", &p2); err != nil {
		panic(err)
	}
	return p1, p2
}

func Part1(lines []string) int {
	p1, p2 := parseStarts(lines)
	scores := [2]int{}
	pos := [2]int{p1, p2}
	die := 0
	rolls := 0

	for {
		for player := 0; player < 2; player++ {
			move := 0
			for i := 0; i < 3; i++ {
				die++
				if die > 100 {
					die = 1
				}
				move += die
				rolls++
			}
			pos[player] = (pos[player]-1+move)%10 + 1
			scores[player] += pos[player]
			if scores[player] >= 1000 {
				loser := 1 - player
				return scores[loser] * rolls
			}
		}
	}
}

// Dirac die: rolling 3 times produces sums 3-9 with these frequencies
var diracRolls = [7][2]int{
	{3, 1}, {4, 3}, {5, 6}, {6, 7}, {7, 6}, {8, 3}, {9, 1},
}

type state struct {
	pos1, pos2, score1, score2 int
	turn                       int // 0 = player 1's turn
}

func Part2(lines []string) int {
	p1, p2 := parseStarts(lines)

	cache := map[state][2]int{}
	var countWins func(s state) [2]int
	countWins = func(s state) [2]int {
		if s.score1 >= 21 {
			return [2]int{1, 0}
		}
		if s.score2 >= 21 {
			return [2]int{0, 1}
		}
		if v, ok := cache[s]; ok {
			return v
		}
		var total [2]int
		for _, roll := range diracRolls {
			sum, freq := roll[0], roll[1]
			ns := s
			if s.turn == 0 {
				ns.pos1 = (s.pos1-1+sum)%10 + 1
				ns.score1 += ns.pos1
			} else {
				ns.pos2 = (s.pos2-1+sum)%10 + 1
				ns.score2 += ns.pos2
			}
			ns.turn = 1 - s.turn
			w := countWins(ns)
			total[0] += w[0] * freq
			total[1] += w[1] * freq
		}
		cache[s] = total
		return total
	}

	wins := countWins(state{pos1: p1, pos2: p2})
	if wins[0] > wins[1] {
		return wins[0]
	}
	return wins[1]
}

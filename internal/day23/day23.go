// Package day23 solves AoC 2021 day 23: Amphipod.
// Find the minimum energy to organize amphipods using Dijkstra over game states.
package day23

import (
	"container/heap"

	"github.com/ajdnik/aoc21/utils"
)

// Amphipod types: 0=A, 1=B, 2=C, 3=D, -1=empty
const empty = -1

var costs = [4]int{1, 10, 100, 1000}
var roomCol = [4]int{2, 4, 6, 8} // hallway positions above rooms

// State encodes hallway (11 slots) + rooms (4 rooms × depth)
type state struct {
	hallway [11]int8
	rooms   [4][]int8
}

func (s state) key() string {
	buf := make([]byte, 0, 11+16)
	for _, h := range s.hallway {
		buf = append(buf, byte(h+2))
	}
	for _, r := range s.rooms {
		for _, v := range r {
			buf = append(buf, byte(v+2))
		}
	}
	return string(buf)
}

func (s state) done() bool {
	for r, room := range s.rooms {
		for _, v := range room {
			if int(v) != r {
				return false
			}
		}
	}
	return true
}

func (s state) clone() state {
	ns := state{hallway: s.hallway}
	for i, room := range s.rooms {
		ns.rooms[i] = make([]int8, len(room))
		copy(ns.rooms[i], room)
	}
	return ns
}

// roomReady: can amphipod type t enter room r?
// Room must be its target and contain only correct types (or empty)
func (s state) roomReady(r int) bool {
	for _, v := range s.rooms[r] {
		if v != empty && int(v) != r {
			return false
		}
	}
	return true
}

// topOfRoom returns index of topmost amphipod in room, or -1
func (s state) topOfRoom(r int) int {
	for i, v := range s.rooms[r] {
		if v != empty {
			return i
		}
	}
	return -1
}

// deepestEmpty returns index of deepest empty slot in room, or -1
func (s state) deepestEmpty(r int) int {
	for i := len(s.rooms[r]) - 1; i >= 0; i-- {
		if s.rooms[r][i] == empty {
			return i
		}
	}
	return -1
}

// hallwayClear checks if hallway is clear between positions a and b (exclusive of a)
func (s state) hallwayClear(from, to int) bool {
	if from < to {
		for i := from + 1; i <= to; i++ {
			if s.hallway[i] != empty {
				return false
			}
		}
	} else {
		for i := from - 1; i >= to; i-- {
			if s.hallway[i] != empty {
				return false
			}
		}
	}
	return true
}

// Valid hallway stopping positions (not directly above rooms)
var hallwayStops = [7]int{0, 1, 3, 5, 7, 9, 10}

type item struct {
	s    state
	cost int
}

type pq []item

func (q pq) Len() int           { return len(q) }
func (q pq) Less(i, j int) bool { return q[i].cost < q[j].cost }
func (q pq) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q *pq) Push(x any)        { *q = append(*q, x.(item)) }
func (q *pq) Pop() any {
	old := *q
	n := len(old)
	v := old[n-1]
	*q = old[:n-1]
	return v
}

func solve(initial state) int {
	dist := map[string]int{}
	q := &pq{{s: initial, cost: 0}}

	for q.Len() > 0 {
		cur := heap.Pop(q).(item)
		if cur.s.done() {
			return cur.cost
		}
		k := cur.s.key()
		if prev, ok := dist[k]; ok && cur.cost >= prev {
			continue
		}
		dist[k] = cur.cost

		// Move from room to hallway
		for r := 0; r < 4; r++ {
			ti := cur.s.topOfRoom(r)
			if ti < 0 {
				continue
			}
			// Skip if room already settled
			if cur.s.roomReady(r) {
				allCorrect := true
				for i := ti; i < len(cur.s.rooms[r]); i++ {
					if int(cur.s.rooms[r][i]) != r {
						allCorrect = false
						break
					}
				}
				if allCorrect {
					continue
				}
			}
			amp := cur.s.rooms[r][ti]
			col := roomCol[r]
			for _, h := range hallwayStops {
				if !cur.s.hallwayClear(col, h) {
					continue
				}
				steps := ti + 1 + utils.Abs(col-h) // steps out of room + hallway distance
				cost := cur.cost + steps*costs[amp]
				ns := cur.s.clone()
				ns.rooms[r][ti] = empty
				ns.hallway[h] = amp
				nk := ns.key()
				if prev, ok := dist[nk]; !ok || cost < prev {
					heap.Push(q, item{s: ns, cost: cost})
				}
			}
		}

		// Move from hallway to room
		for h := 0; h < 11; h++ {
			amp := cur.s.hallway[h]
			if amp == empty {
				continue
			}
			r := int(amp)
			if !cur.s.roomReady(r) {
				continue
			}
			col := roomCol[r]
			if !cur.s.hallwayClear(h, col) {
				continue
			}
			di := cur.s.deepestEmpty(r)
			if di < 0 {
				continue
			}
			steps := utils.Abs(h-col) + di + 1
			cost := cur.cost + steps*costs[amp]
			ns := cur.s.clone()
			ns.hallway[h] = empty
			ns.rooms[r][di] = amp
			nk := ns.key()
			if prev, ok := dist[nk]; !ok || cost < prev {
				heap.Push(q, item{s: ns, cost: cost})
			}
		}
	}
	return -1
}

func parseInput(lines []string, depth int) state {
	s := state{}
	for i := range s.hallway {
		s.hallway[i] = empty
	}
	for r := 0; r < 4; r++ {
		s.rooms[r] = make([]int8, depth)
	}
	for d := 0; d < depth; d++ {
		row := lines[2+d]
		for r := 0; r < 4; r++ {
			ch := row[3+r*2]
			s.rooms[r][d] = int8(ch - 'A')
		}
	}
	return s
}

func Part1(lines []string) int {
	return solve(parseInput(lines, 2))
}

func Part2(lines []string) int {
	// Insert two extra rows
	extended := make([]string, 0, len(lines)+2)
	extended = append(extended, lines[:3]...)
	extended = append(extended, "  #D#C#B#A#", "  #D#B#A#C#")
	extended = append(extended, lines[3:]...)
	return solve(parseInput(extended, 4))
}

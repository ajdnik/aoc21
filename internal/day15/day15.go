package day15

import (
	"container/heap"
)

type point struct {
	x, y int
}

type item struct {
	pos  point
	risk int
}

type priorityQueue []item

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].risk < pq[j].risk }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(item)) }
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	v := old[n-1]
	*pq = old[:n-1]
	return v
}

var dirs = [4]point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func dijkstra(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	dist := make([][]int, rows)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := range dist[i] {
			dist[i][j] = 1<<63 - 1
		}
	}
	dist[0][0] = 0

	pq := &priorityQueue{{pos: point{0, 0}, risk: 0}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(item)
		if cur.pos.x == rows-1 && cur.pos.y == cols-1 {
			return cur.risk
		}
		if cur.risk > dist[cur.pos.x][cur.pos.y] {
			continue
		}
		for _, d := range dirs {
			nx, ny := cur.pos.x+d.x, cur.pos.y+d.y
			if nx < 0 || nx >= rows || ny < 0 || ny >= cols {
				continue
			}
			newRisk := cur.risk + grid[nx][ny]
			if newRisk < dist[nx][ny] {
				dist[nx][ny] = newRisk
				heap.Push(pq, item{pos: point{nx, ny}, risk: newRisk})
			}
		}
	}
	return dist[rows-1][cols-1]
}

func parseGrid(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, ch := range line {
			grid[i][j] = int(ch - '0')
		}
	}
	return grid
}

func tileGrid(grid [][]int, factor int) [][]int {
	rows := len(grid)
	cols := len(grid[0])
	tiled := make([][]int, rows*factor)
	for i := range tiled {
		tiled[i] = make([]int, cols*factor)
		for j := range tiled[i] {
			val := grid[i%rows][j%cols] + i/rows + j/cols
			if val > 9 {
				val -= 9
			}
			tiled[i][j] = val
		}
	}
	return tiled
}

func Part1(lines []string) int {
	return dijkstra(parseGrid(lines))
}

func Part2(lines []string) int {
	return dijkstra(tileGrid(parseGrid(lines), 5))
}

package solution

import (
	"math"
)

func expandMatrix(m [][]int) [][]int {
	bm := make([][]int, len(m)*5)
	for i := 0; i < len(bm); i++ {
		bm[i] = make([]int, len(m[0])*5)
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			for ii := 0; ii < 5; ii++ {
				for jj := 0; jj < 5; jj++ {
					val := (m[i][j] + ii + jj)
					if val > 9 {
						val = val - 9
					}
					bm[i+len(m)*ii][j+len(m[0])*jj] = val
				}
			}
		}
	}
	return bm
}

func dijkstra(m [][]int) int {
	distance := make(map[Point]int)
	unvisited := make(map[Point]bool)
	candidate := make(map[Point]bool)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			distance[Point{i, j}] = math.MaxInt
			unvisited[Point{i, j}] = true
		}
	}

	visit := func(p Point, currentDistance int) {
		if p.x < 0 || p.x >= len(m) {
			return
		}
		if p.y < 0 || p.y >= len(m[0]) {
			return
		}
		if !unvisited[p] {
			return
		}
		distance[p] = min([]int{distance[p], currentDistance + m[p.x][p.y]})
		candidate[p] = true
	}

	current := Point{0, 0}

	distance[current] = 0
	for {
		delete(unvisited, current)
		delete(candidate, current)
		visit(Point{current.x - 1, current.y}, distance[current])
		visit(Point{current.x + 1, current.y}, distance[current])
		visit(Point{current.x, current.y - 1}, distance[current])
		visit(Point{current.x, current.y + 1}, distance[current])
		next := Point{}
		nextDistance := math.MaxInt
		for p, _ := range candidate {
			if distance[p] < nextDistance {
				next = p
				nextDistance = distance[p]
			}
		}
		if next.x == len(m)-1 && next.y == len(m[0])-1 {
			return distance[next]
		}
		current = next
	}
}

func (s *Solution) Day15Part1(fn string) (ret int) {
	m := toMatrix(fn)

	return dijkstra(m)
}

func (s *Solution) Day15Part2(fn string) (ret int) {
	m := toMatrix(fn)
	m = expandMatrix(m)

	return dijkstra(m)
}

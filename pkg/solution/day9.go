package solution

import (
	"sort"
	"strings"
)

func (s *Solution) Day9Part1(fn string) (ret int) {
	lines := toLines(fn)

	m := [][]int{}
	for _, l := range lines {
		split := strings.Split(l, "")
		ints := toInts(split)
		m = append(m, ints)
	}

	cnt := 0
	numRow := len(m)
	numColumn := len(m[0])
	for i := 0; i < numRow; i++ {
		for j := 0; j < numColumn; j++ {
			// check if matrix[i][j] is low point
			if (j == 0 || m[i][j-1] > m[i][j]) &&
				(j == (numColumn-1) || m[i][j+1] > m[i][j]) &&
				(i == 0 || m[i-1][j] > m[i][j]) &&
				(i == (numRow-1) || m[i+1][j] > m[i][j]) {
				cnt = cnt + m[i][j] + 1
			}
		}
	}

	return cnt
}

func floodFill(p Point, visited map[Point]bool, m [][]int, maxX int, maxY int) []Point {
	ret := []Point{p}
	visited[p] = true
	// left
	if p.y > 0 && m[p.x][p.y-1] > m[p.x][p.y] && m[p.x][p.y-1] != 9 && !visited[Point{p.x, p.y - 1}] {
		filled := floodFill(Point{p.x, p.y - 1}, visited, m, maxX, maxY)
		ret = append(ret, filled...)
	}
	// right
	if p.y < maxY-1 && m[p.x][p.y+1] > m[p.x][p.y] && m[p.x][p.y+1] != 9 && !visited[Point{p.x, p.y + 1}] {
		filled := floodFill(Point{p.x, p.y + 1}, visited, m, maxX, maxY)
		ret = append(ret, filled...)
	}
	// up
	if p.x > 0 && m[p.x-1][p.y] > m[p.x][p.y] && m[p.x-1][p.y] != 9 && !visited[Point{p.x - 1, p.y}] {
		filled := floodFill(Point{p.x - 1, p.y}, visited, m, maxX, maxY)
		ret = append(ret, filled...)
	}
	// down
	if p.x < maxX-1 && m[p.x+1][p.y] > m[p.x][p.y] && m[p.x+1][p.y] != 9 && !visited[Point{p.x + 1, p.y}] {
		filled := floodFill(Point{p.x + 1, p.y}, visited, m, maxX, maxY)
		ret = append(ret, filled...)
	}

	return ret
}

func (s *Solution) Day9Part2(fn string) (ret int) {
	lines := toLines(fn)

	m := [][]int{}
	for _, l := range lines {
		split := strings.Split(l, "")
		ints := toInts(split)
		m = append(m, ints)
	}

	visited := make(map[Point]bool)
	sizes := []int{}
	numRow := len(m)
	numColumn := len(m[0])
	for i := 0; i < numRow; i++ {
		for j := 0; j < numColumn; j++ {
			// check if matrix[i][j] is low point
			if (j == 0 || m[i][j-1] > m[i][j]) &&
				(j == (numColumn-1) || m[i][j+1] > m[i][j]) &&
				(i == 0 || m[i-1][j] > m[i][j]) &&
				(i == (numRow-1) || m[i+1][j] > m[i][j]) {
				// is low point
				basin := floodFill(Point{i, j}, visited, m, numRow, numColumn)
				for _, p := range basin {
					visited[p] = true
				}
				sizes = append(sizes, len(basin))
			}
		}
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[j] < sizes[i]
	})
	return sizes[0] * sizes[1] * sizes[2]
}

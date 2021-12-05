package solution

import (
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y int
}

func parseCoordinates(s string) (l Coordinate, r Coordinate) {
	p := strings.Split(s, "->")
	p1 := strings.Split(p[0], ",")
	p2 := strings.Split(p[1], ",")

	lx, _ := strconv.Atoi(p1[0])
	ly, _ := strconv.Atoi(p1[1])

	rx, _ := strconv.Atoi(p2[0])
	ry, _ := strconv.Atoi(p2[1])

	return Coordinate{x: lx, y: ly}, Coordinate{x: rx, y: ry}
}

func (s *Solution) Day5Part1(fn string) (ret int) {
	lines := toLines(fn)
	m := make(map[Coordinate]int)
	cnt := 0

	for _, s := range lines {
		s = strings.ReplaceAll(s, " ", "")
		c1, c2 := parseCoordinates(s)

		if c1.x == c2.x {
			y1, y2 := sort(c1.y, c2.y)
			for i := y1; i <= y2; i++ {
				coordCnt := m[Coordinate{c1.x, i}]
				if coordCnt == 1 {
					cnt += 1
				}
				m[Coordinate{c1.x, i}] = coordCnt + 1
			}
		}

		if c1.y == c2.y {
			x1, x2 := sort(c1.x, c2.x)
			for i := x1; i <= x2; i++ {
				coordCnt := m[Coordinate{i, c1.y}]
				if coordCnt == 1 {
					cnt += 1
				}
				m[Coordinate{i, c1.y}] = coordCnt + 1
			}
		}
	}
	return cnt
}

func (s *Solution) Day5Part2(fn string) (ret int) {
	lines := toLines(fn)
	m := make(map[Coordinate]int)
	cnt := 0

	for _, s := range lines {
		s = strings.ReplaceAll(s, " ", "")
		c1, c2 := parseCoordinates(s)

		dx, dy := 0, 0
		if c1.x < c2.x {
			dx = 1
		} else if c1.x > c2.x {
			dx = -1
		}

		if c1.y < c2.y {
			dy = 1
		} else if c1.y > c2.y {
			dy = -1
		}

		for {
			m[c1] += 1
			if c1.x == c2.x && c1.y == c2.y {
				break
			}
			c1.x += dx
			c1.y += dy
		}
	}
	for _, v := range m {
		if v > 1 {
			cnt += 1
		}
	}
	return cnt
}

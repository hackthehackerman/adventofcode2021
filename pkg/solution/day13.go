package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func (s *Solution) Day13Part1(fn string) (ret int) {
	lines := toLines(fn)

	foldStartIdx := 0
	maxX := 0
	maxY := 0
	dots := make(map[Point]bool)
	for i, l := range lines {
		if l == "" {
			foldStartIdx = i + 1
			break
		}

		split := strings.Split(l, ",")
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])

		maxX = max([]int{maxX, l})
		maxY = max([]int{maxY, r})
		dots[Point{l, r}] = true
	}

	for i := foldStartIdx; i < foldStartIdx+1; i++ {
		split := strings.Fields(lines[i])
		split = strings.Split(split[2], "=")

		l := split[0]
		n, _ := strconv.Atoi(split[1])
		if l == "x" {
			for x := n + 1; x <= maxX; x++ {
				for y := 0; y <= maxY; y++ {
					if dots[Point{x, y}] {
						delete(dots, Point{x, y})
						dots[Point{n - (x - n), y}] = true
					}
				}
			}
		}

		if l == "y" {
			for y := n + 1; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {
					if dots[Point{x, y}] {
						delete(dots, Point{x, y})
						dots[Point{x, n - (y - n)}] = true
					}
				}
			}
		}
	}

	return len(dots)
}

func (s *Solution) Day13Part2(fn string) int {
	lines := toLines(fn)

	foldStartIdx := 0
	maxX := 0
	maxY := 0
	dots := make(map[Point]bool)
	for i, l := range lines {
		if l == "" {
			foldStartIdx = i + 1
			break
		}

		split := strings.Split(l, ",")
		l, _ := strconv.Atoi(split[0])
		r, _ := strconv.Atoi(split[1])

		maxX = max([]int{maxX, l})
		maxY = max([]int{maxY, r})
		dots[Point{l, r}] = true
	}

	for i := foldStartIdx; i < len(lines); i++ {
		split := strings.Fields(lines[i])
		split = strings.Split(split[2], "=")

		l := split[0]
		n, _ := strconv.Atoi(split[1])
		if l == "x" {
			for x := n + 1; x <= maxX; x++ {
				for y := 0; y <= maxY; y++ {
					if dots[Point{x, y}] {
						delete(dots, Point{x, y})
						dots[Point{n - (x - n), y}] = true
					}
				}
			}
		}

		if l == "y" {
			for y := n + 1; y <= maxY; y++ {
				for x := 0; x <= maxX; x++ {
					if dots[Point{x, y}] {
						delete(dots, Point{x, y})
						dots[Point{x, n - (y - n)}] = true
					}
				}
			}
		}
	}

	// render
	maxX = 0
	maxY = 0
	for k := range dots {
		maxX = max([]int{maxX, k.x})
		maxY = max([]int{maxY, k.y})
	}

	m := [][]string{}
	for x := 0; x <= maxX; x++ {
		row := []string{}
		for y := 0; y <= maxY; y++ {
			if dots[Point{x, y}] {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
		}
		m = append(m, row)
	}

	for i := len(m) - 1; i >= 0; i-- {
		fmt.Println(m[i])
	}

	return 0
}

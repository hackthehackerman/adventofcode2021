package solution

import (
	"math"
	"strings"
)

func buildMap(n []int) (m map[int]int) {
	m = make(map[int]int)
	for i, num := range n {
		m[num] = i
	}
	return
}

func reverse(m map[int]int) (rm map[int]int) {
	rm = make(map[int]int)
	for k, v := range m {
		rm[v] = k
	}
	return
}

func next(l []string, idx int) (board [][]int, nextIdx int) {
	if idx >= len(l) {
		return
	}

	for idx < len(l) && l[idx] != "" {
		row := strings.Fields(l[idx])
		rowInt := toInts(row)
		board = append(board, rowInt)
		idx++
	}
	return board, idx + 1
}

func column(b [][]int, c int) (l []int) {
	for i := 0; i < len(b); i++ {
		l = append(l, b[i][c])
	}
	return
}

func evalLine(l []int, m map[int]int) (step int) {
	step = 0
	for _, n := range l {
		if val, ok := m[n]; ok && val > step {
			step = val
		}
	}
	return
}

func eval(b [][]int, m map[int]int) (step int, score int) {
	step = math.MaxInt
	for i := 0; i < len(b); i++ {
		var s int
		s = evalLine(b[i], m)
		if s < step {
			step = s
		}
		s = evalLine(column(b, i), m)
		if s < step {
			step = s
		}
	}

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			n := b[i][j]
			if val, ok := m[n]; ok && val > step {
				score += n
			}
		}
	}
	return
}

func (s *Solution) Day4Part1(fn string) (ret int) {
	lines := toLines(fn)

	split := strings.Split(lines[0], ",")
	n := toInts(split)
	m := buildMap(n)
	rm := reverse(m)

	var step int = math.MaxInt
	var score int = math.MaxInt

	b, i := next(lines, 2)
	for len(b) != 0 {
		stepTmp, scoreTmp := eval(b, m)
		if stepTmp < step {
			step = stepTmp
			score = scoreTmp
		}

		b, i = next(lines, i)
	}

	return rm[step] * score
}

func (s *Solution) Day4Part2(fn string) (ret int) {
	lines := toLines(fn)

	split := strings.Split(lines[0], ",")
	n := toInts(split)
	m := buildMap(n)
	rm := reverse(m)

	var step int = 0
	var score int = math.MaxInt

	b, i := next(lines, 2)
	for len(b) != 0 {
		stepTmp, scoreTmp := eval(b, m)
		if stepTmp > step {
			step = stepTmp
			score = scoreTmp
		}

		b, i = next(lines, i)
	}

	return rm[step] * score
}

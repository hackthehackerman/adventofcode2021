package solution

import (
	"strings"
)

func (s *Solution) Day8Part1(fn string) (ret int) {
	lines := toLines(fn)
	cnt := 0
	for _, l := range lines {
		signals := strings.Split(l, "|")
		signals_out := strings.Fields(signals[1])

		for _, s := range signals_out {
			switch len(s) {
			case 2, 3, 4, 7:
				cnt += 1
			}
		}
	}

	return cnt
}

type Board struct {
	combinationMap map[string]int
}

func generateBoards() []Board {
	makeCombination := func(b []string, pos []int) string {
		var sb strings.Builder
		for _, i := range pos {
			sb.WriteString(b[i])
		}
		return sortString(sb.String())
	}

	permutations := Perm([]string{"a", "b", "c", "d", "e", "f", "g"})
	boards := []Board{}
	for _, v := range permutations {
		m := make(map[string]int)
		m[makeCombination(v, []int{0, 1, 2, 4, 5, 6})] = 0
		m[makeCombination(v, []int{2, 5})] = 1
		m[makeCombination(v, []int{0, 2, 3, 4, 6})] = 2
		m[makeCombination(v, []int{0, 2, 3, 5, 6})] = 3
		m[makeCombination(v, []int{1, 2, 3, 5})] = 4
		m[makeCombination(v, []int{0, 1, 3, 5, 6})] = 5
		m[makeCombination(v, []int{0, 1, 3, 4, 5, 6})] = 6
		m[makeCombination(v, []int{0, 2, 5})] = 7
		m[makeCombination(v, []int{0, 1, 2, 3, 4, 5, 6})] = 8
		m[makeCombination(v, []int{0, 1, 2, 3, 5, 6})] = 9
		boards = append(boards, Board{m})
	}
	return boards
}

func boardFor(boards []Board, in []string) Board {
OUTER:
	for _, b := range boards {
		for _, str := range in {
			sorted := sortString(str)
			if _, ok := b.combinationMap[sorted]; !ok {
				continue OUTER
			}
		}
		return b
	}
	panic("shouldn't be here")
}

func (s *Solution) Day8Part2(fn string) (ret int) {
	lines := toLines(fn)
	cnt := 0

	// generate all possible board
	boards := generateBoards()
	for _, l := range lines {
		signals := strings.Split(l, "|")
		signalsIn := strings.Fields(signals[0])
		signalsOut := strings.Fields(signals[1])

		b := boardFor(boards, signalsIn)
		out := 0
		for _, o := range signalsOut {
			sorted := sortString(o)
			out = out*10 + b.combinationMap[sorted]
		}
		cnt += out
	}

	return cnt
}

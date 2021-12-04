package solution

import (
	"strconv"
)

type Solution struct {
}

func (s *Solution) Day1Part1(fn string) (ret int) {
	lines := toLines(fn)
	prev := 0
	cnt := 0
	for i, s := range lines {
		n, _ := strconv.Atoi(s)
		if i > 0 && n > prev {
			cnt += 1
		}
		prev = n
	}
	return cnt
}

func (s *Solution) Day1Part2(fn string) (ret int) {
	nums := toInts(toLines(fn))
	cnt := 0
	for i, n := range nums {
		if i > 2 && n > nums[i-3] {
			cnt += 1
		}
	}
	return cnt
}

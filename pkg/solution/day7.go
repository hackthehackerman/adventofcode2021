package solution

import (
	"math"
	"strings"
)

func delta(i int, ints []int) int {
	ret := 0
	for _, v := range ints {
		ret += int(math.Abs(float64(v - i)))
	}
	return ret
}

func delta2(i int, ints []int) int {
	ret := 0
	for _, v := range ints {
		distance := int(math.Abs(float64(v - i)))
		sumDistance := ((1 + distance) * distance) / 2
		ret += sumDistance
	}
	return ret
}

func (s *Solution) Day7Part1(fn string) (ret int) {
	lines := toLines(fn)
	nums := strings.Split(lines[0], ",")
	numsInts := toInts(nums)

	ret = math.MaxInt
	retDelta := math.MaxInt
	for i := min(numsInts); i <= max(numsInts); i++ {
		d := delta(i, numsInts)
		if d < retDelta {
			ret = i
			retDelta = d
		}
	}

	return retDelta
}

func (s *Solution) Day7Part2(fn string) (ret int) {
	lines := toLines(fn)
	nums := strings.Split(lines[0], ",")
	numsInts := toInts(nums)

	ret = math.MaxInt
	retDelta := math.MaxInt
	for i := min(numsInts); i <= max(numsInts); i++ {
		d := delta2(i, numsInts)
		if d < retDelta {
			ret = i
			retDelta = d
		}
	}

	return retDelta
}

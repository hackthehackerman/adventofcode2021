package solution

import (
	"strings"
)

type Pair struct {
	fish, day int
}

func simulateDay(fish []int) []int {
	numNewFish := 0
	for j, f := range fish {
		if f == 0 {
			numNewFish++
			fish[j] = 6
		} else {
			fish[j] = f - 1
		}
	}
	for j := 0; j < numNewFish; j++ {
		fish = append(fish, 8)
	}
	return fish
}

func allCached(fish []int, day int, cache map[Pair]int) bool {
	for _, f := range fish {
		if _, ok := cache[Pair{f, day}]; !ok {
			return false
		}
	}
	return true
}

func (s *Solution) calc(start []int, maxDay int) (ret int) {
	maxFish := 8
	cache := make(map[Pair]int)

	// pre compute number of fish with initial fish num and number of days
	fish := []int{maxFish}
	for i := 0; i < (maxDay/2)+maxFish+1; i++ {
		numFish := len(fish)
		for j := maxFish; j >= 0; j-- {
			d := i - (maxFish - j)
			if d >= 0 {
				cache[Pair{j, d}] = numFish
			}
		}
		fish = simulateDay(fish)
	}

	fish = start
	cnt := 0
	for i := 0; i < maxDay/2+1; i++ {
		dayRemained := maxDay - i
		if allCached(fish, dayRemained, cache) {
			for _, f := range fish {
				cnt += cache[Pair{f, dayRemained}]
			}
			return cnt
		} else {
			fish = simulateDay(fish)
		}
	}
	return 0
}

func (s *Solution) Day6Part1(fn string) (ret int) {
	lines := toLines(fn)
	nums := strings.Split(lines[0], ",")
	numsInts := toInts(nums)

	return s.calc(numsInts, 80)
}

func (s *Solution) Day6Part2(fn string) (ret int) {
	lines := toLines(fn)
	nums := strings.Split(lines[0], ",")
	numsInts := toInts(nums)

	return s.calc(numsInts, 256)
}

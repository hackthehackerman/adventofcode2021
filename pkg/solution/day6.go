package solution

import (
	"fmt"
	"strings"
)

func (s *Solution) Day6Part1(fn string) (ret int) {
	lines := toLines(fn)
	nums := strings.Split(lines[0], ",")
	numsInts := toInts(nums)

	// simulate 1 fish with initial number 6
	fish := []int{6}
	cnt := make(map[int]int)
	for i := 0; i < 86; i++ {
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
		cnt[i] = len(fish)
	}

	totalCnt := 0
	for _, n := range numsInts {
		totalCnt = totalCnt + cnt[85-n]
	}

	return totalCnt
}

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

func calc(fish int, curDay int, maxDay int, cache map[Pair]int) int {
	// time.Sleep(1 * time.Second)
	// fmt.Println(fish, curDay, maxDay, cache)

	if val, ok := cache[Pair{fish, maxDay - curDay}]; ok {
		return val
	}
	if curDay == maxDay {
		return 1
	}

	fishArr := []int{fish}
	fishArr = simulateDay(fishArr)
	// cache[Pair{fish, maxDay - curDay - 1}] = len(fishArr)
	for i := fish; i <= 6; i++ {

		cache[Pair{i, maxDay - curDay - 1 + (i - fish)}] = len(fishArr)
	}

	cnt := 0
	for _, f := range fishArr {
		cnt += calc(f, curDay+1, maxDay, cache)
	}
	return cnt
}

func (s *Solution) Day6Part2(fn string) (ret int) {
	lines := toLines(fn)
	nums := strings.Split(lines[0], ",")
	numsInts := toInts(nums)

	maxDay := 256
	maxFish := 8
	cache := make(map[Pair]int)

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

	fish = numsInts
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

	fmt.Println(cache)

	return 0
}

package solution

import (
	"strings"
)

func (s *Solution) Day14Part1(fn string) (ret int) {
	lines := toLines(fn)

	start := lines[0]
	rules := make(map[string]string)

	for i := 2; i < len(lines); i++ {
		split := strings.Fields(lines[i])
		rules[split[0]] = split[2]
	}

	for i := 0; i < 10; i++ {
		var sb strings.Builder
		for j := 0; j < len(start); j++ {
			if j == len(start)-1 {
				sb.WriteString(start[j : j+1])
			} else if val, ok := rules[start[j:j+2]]; ok {
				sb.WriteString(start[j : j+1])
				sb.WriteString(val)
			} else {
				sb.WriteString(start[j : j+1])
			}
		}
		start = sb.String()
	}

	counter := make(map[rune]int)
	for _, r := range start {
		counter[r] += 1
	}

	nums := []int{}
	for _, v := range counter {
		nums = append(nums, v)
	}

	return max(nums) - min(nums)
}

func (s *Solution) Day14Part2(fn string) (ret int) {
	lines := toLines(fn)

	start := lines[0]
	rules := make(map[string]string)
	for i := 2; i < len(lines); i++ {
		split := strings.Fields(lines[i])
		rules[split[0]] = split[2]
	}

	pairs := make(map[string]int)
	for j := 0; j < len(start)-1; j++ {
		pairs[start[j:j+2]] += 1
	}
	counter := make(map[string]int)
	for j := 0; j < len(start); j++ {
		counter[start[j:j+1]] += 1
	}

	for i := 0; i < 40; i++ {
		newPairs := make(map[string]int)
		for k, v := range pairs {
			if val, ok := rules[k]; ok {
				lp := k[0:1] + val
				rp := val + k[1:2]
				newPairs[lp] += v
				newPairs[rp] += v
				pairs[k] -= v
				counter[val] += v
			}
		}
		for k, v := range newPairs {
			pairs[k] += v
		}
	}

	nums := []int{}
	for _, v := range counter {
		nums = append(nums, v)
	}

	return max(nums) - min(nums)
}

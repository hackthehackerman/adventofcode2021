package solution

import (
	"sort"
)

func (s *Solution) Day10Part1(fn string) (ret int) {
	lines := toLines(fn)
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	cnt := 0
OUTER:
	for _, l := range lines {
		stack := []rune{}
		for _, r := range l {
			if _, ok := pairs[r]; ok {
				stack = append(stack, r)
			} else {
				if len(stack) == 0 ||
					pairs[stack[len(stack)-1]] != r {
					// invalid closing
					cnt += scores[r]
					continue OUTER
				} else {
					// valid closing
					stack = stack[:len(stack)-1]
				}
			}
		}
	}

	return cnt
}

func (s *Solution) Day10Part2(fn string) (ret int) {
	lines := toLines(fn)
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	scores := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	scoresArr := []int{}
OUTER:
	for _, l := range lines {
		stack := []rune{}
		for i, r := range l {
			if _, ok := pairs[r]; ok {
				stack = append(stack, r)
			} else {
				if len(stack) == 0 ||
					pairs[stack[len(stack)-1]] != r {
					// invalid closing
					continue OUTER
				} else {
					// valid closing
					stack = stack[:len(stack)-1]
				}
			}
			if i == len(l)-1 {
				// incomplete
				s := 0
				for j := len(stack) - 1; j >= 0; j-- {
					s = s*5 + scores[pairs[stack[j]]]
				}
				scoresArr = append(scoresArr, s)
			}
		}
	}

	sort.Ints(scoresArr)

	return scoresArr[len(scoresArr)/2]
}

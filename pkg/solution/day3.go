package solution

import (
	"strconv"
	"strings"
)

func (s *Solution) Day3Part1(fn string) (ret int64) {
	lines := toLines("f3_1")
	var bits [12]int
	for _, s := range lines {
		for i, c := range s {
			switch c {
			case '0':
				bits[i] -= 1
			case '1':
				bits[i] += 1
			}
		}
	}

	var sb_g strings.Builder
	var sb_e strings.Builder
	for _, n := range bits {
		if n > 0 {
			sb_g.WriteString("1")
			sb_e.WriteString("0")
		} else {
			sb_g.WriteString("0")
			sb_e.WriteString("1")
		}
	}

	gamma, _ := strconv.ParseInt(sb_g.String(), 2, 64)
	epsilon, _ := strconv.ParseInt(sb_e.String(), 2, 64)
	return gamma * epsilon
}

func (s *Solution) Day3Part2(fn string) (ret int64) {
	lines := toLines(fn)

	var getReading func(flag string, lines []string, idx int) int64
	getReading = func(flag string, lines []string, idx int) int64 {
		if len(lines) == 0 {
			return 0
		}
		if len(lines) == 1 {
			ret, _ := strconv.ParseInt(lines[0], 2, 64)
			return ret
		}
		m := map[string][]string{"0": {}, "1": {}}
		cnt := 0
		for _, s := range lines {
			m[string(s[idx])] = append(m[string(s[idx])], s)
			switch string(s[idx]) {
			case "0":
				cnt -= 1
			case "1":
				cnt += 1
			}
		}

		if cnt >= 0 {
			return getReading(flag, m[flag], idx+1)
		} else if flag == "0" {
			return getReading(flag, m["1"], idx+1)
		} else {
			return getReading(flag, m["0"], idx+1)
		}
	}

	o := getReading("1", lines, 0)
	c := getReading("0", lines, 0)

	return o * c
}

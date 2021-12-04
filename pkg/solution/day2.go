package solution

import (
	"strconv"
	"strings"
)

func (s *Solution) Day2Part1(fn string) (ret int) {
	lines := toLines(fn)
	h := 0
	v := 0
	for _, s := range lines {
		split := strings.Split(s, " ")
		dir := split[0]
		dist, _ := strconv.Atoi(split[1])
		switch dir {
		case "forward":
			h += dist
		case "up":
			v -= dist
		case "down":
			v += dist
		}
	}
	return h * v
}

func (s *Solution) Day2Part2(fn string) (ret int) {
	lines := toLines(fn)
	h := 0
	d := 0
	a := 0
	for _, s := range lines {
		split := strings.Split(s, " ")
		dir := split[0]
		dist, _ := strconv.Atoi(split[1])
		switch dir {
		case "forward":
			h += dist
			d += dist * a
		case "up":
			a -= dist
		case "down":
			a += dist
		}
	}
	return h * d
}

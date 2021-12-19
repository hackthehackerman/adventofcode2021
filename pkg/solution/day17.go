package solution

import (
	"math"
	"strconv"
	"strings"
)

func (s *Solution) Day17Part1(fn string) (ret int) {
	l := toLines(fn)[0]
	split := strings.Fields(l)
	x := strings.ReplaceAll(split[2], "x=", "")
	x = strings.ReplaceAll(x, ",", "")
	y := strings.ReplaceAll(split[3], "y=", "")
	xs := strings.Split(x, "..")
	ys := strings.Split(y, "..")

	xl, _ := strconv.Atoi(xs[0])
	xr, _ := strconv.Atoi(xs[1])
	yl, _ := strconv.Atoi(ys[0])
	yr, _ := strconv.Atoi(ys[1])

	simulate := func(x, y, xl, xr, yl, yr int) (maxY, step int, hit bool) {
		posX, posY := 0, 0
		maxY = math.MinInt
		for i := 1; ; i++ {
			posX += x
			posY += y
			if posX > xr || posY < yl {
				break
			}
			if posX >= xl && posX <= xr && posY >= yl && posY <= yr {
				hit = true
			}
			if posY > maxY {
				maxY = posY
			}
			if x > 0 {
				x -= 1
			}
			y -= 1
		}
		return
	}

	maxY := math.MinInt
	for x := 1; x <= xr; x++ {
		for y := yl; y <= int(math.Abs(float64(yl))); y++ {
			my, _, hit := simulate(x, y, xl, xr, yl, yr)
			if hit && my > maxY {
				maxY = my
			}
		}

	}

	return maxY
}

func (s *Solution) Day17Part2(fn string) (ret int) {
	l := toLines(fn)[0]
	split := strings.Fields(l)
	x := strings.ReplaceAll(split[2], "x=", "")
	x = strings.ReplaceAll(x, ",", "")
	y := strings.ReplaceAll(split[3], "y=", "")
	xs := strings.Split(x, "..")
	ys := strings.Split(y, "..")

	xl, _ := strconv.Atoi(xs[0])
	xr, _ := strconv.Atoi(xs[1])
	yl, _ := strconv.Atoi(ys[0])
	yr, _ := strconv.Atoi(ys[1])

	simulate := func(x, y, xl, xr, yl, yr int) (maxY, step int, hit bool) {
		posX, posY := 0, 0
		maxY = math.MinInt
		for i := 1; ; i++ {
			posX += x
			posY += y
			if posX > xr || posY < yl {
				break
			}
			if posX >= xl && posX <= xr && posY >= yl && posY <= yr {
				hit = true
			}
			if posY > maxY {
				maxY = posY
			}
			if x > 0 {
				x -= 1
			}
			y -= 1
		}
		return
	}

	cnt := 0
	for x := 1; x <= xr; x++ {
		for y := yl; y <= int(math.Abs(float64(yl))); y++ {
			_, _, hit := simulate(x, y, xl, xr, yl, yr)
			if hit {
				cnt += 1
			}
		}

	}

	return cnt
}

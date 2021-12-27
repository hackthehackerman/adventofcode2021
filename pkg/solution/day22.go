package solution

import (
	"regexp"
	"strconv"
	"strings"
)

func (s *Solution) Day22Part1(fn string) (ret int) {
	lines := toLines(fn)

	getRange := func(s string) (l, r int, overlapped bool) {
		sl, sr := strings.Split(s, "..")[0], strings.Split(s, "..")[1]
		l, _ = strconv.Atoi(sl)
		r, _ = strconv.Atoi(sr)

		if l > 50 || r < -50 {
			l, r = 0, 0
			return
		}

		if l < -50 {
			l = -50
		} else if l > 50 {
			l = 50
		}

		if r < -50 {
			r = -50
		} else if r > 50 {
			r = 50
		}

		return l, r, true
	}

	type Point3D struct {
		x, y, z int
	}

	onoffs := map[Point3D]bool{}
	for _, l := range lines {
		op := strings.Fields(l)[0]

		r, _ := regexp.Compile("x=(.*?),y=(.*?),z=(.*)")
		f := r.FindAllStringSubmatch(l, -1)

		xRange, yRange, zRange := f[0][1], f[0][2], f[0][3]

		xl, xr, xo := getRange(xRange)
		yl, yr, yo := getRange(yRange)
		zl, zr, zo := getRange(zRange)

		if !xo || !yo || !zo {
			continue
		}

		for x := xl; x <= xr; x++ {
			for y := yl; y <= yr; y++ {
				for z := zl; z <= zr; z++ {
					if op == "on" {
						onoffs[Point3D{x, y, z}] = true
					} else {
						delete(onoffs, Point3D{x, y, z})
					}
				}
			}
		}
	}
	return len(onoffs)
}

func (s *Solution) Day22Part2(fn string) (ret int) {
	type Cube struct {
		xl, xr, yl, yr, zl, zr int
	}

	lines := toLines(fn)

	getRange := func(s string) (l, r int) {
		sl, sr := strings.Split(s, "..")[0], strings.Split(s, "..")[1]
		l, _ = strconv.Atoi(sl)
		r, _ = strconv.Atoi(sr)

		return
	}

	inrange := func(p, l, r int) bool {
		if p >= l && p <= r {
			return true
		}
		return false
	}

	overlappedRange := func(l1, r1, l2, r2 int) (l, r int, overlapped bool) {
		l = max([]int{l1, l2})
		r = min([]int{r1, r2})
		if l > r {
			return 0, 0, false
		} else {
			return l, r, true
		}
	}

	// assume l2, r2 overlapped with l1,r2, and doesn't cover l1,r1 entirely
	split := func(l1, r1, l2, r2 int) (ret []int) {
		tmp := []int{}
		if inrange(l2, l1, r1) && inrange(r2, l1, r1) {
			tmp = []int{l1, l2 - 1, l2, r2, r2 + 1, r1}
		} else if inrange(r2, l1, r1) {
			tmp = []int{l1, r2, r2 + 1, r1}
		} else if inrange(l2, l1, r1) {
			tmp = []int{l1, l2 - 1, l2, r1}
		}

		for i := 0; i < len(tmp); i += 2 {
			if tmp[i+1]-tmp[i] < 0 {
				continue
			} else {
				ret = append(ret, tmp[i])
				ret = append(ret, tmp[i+1])
			}
		}
		return
	}

	overlappedCube := func(c1, c2 Cube) (ret Cube, overlapped bool) {
		xl, xr, xo := overlappedRange(c1.xl, c1.xr, c2.xl, c2.xr)
		yl, yr, yo := overlappedRange(c1.yl, c1.yr, c2.yl, c2.yr)
		zl, zr, zo := overlappedRange(c1.zl, c1.zr, c2.zl, c2.zr)

		if !xo || !yo || !zo {
			return Cube{}, false
		} else {
			return Cube{xl, xr, yl, yr, zl, zr}, true
		}
	}

	subtractCubes := func(c1, c2 Cube) (ret []Cube) {
		c, o := overlappedCube(c1, c2)

		if !o {
			return []Cube{c1}
		}

		xsplit := split(c1.xl, c1.xr, c.xl, c.xr)
		ysplit := split(c1.yl, c1.yr, c.yl, c.yr)
		zsplit := split(c1.zl, c1.zr, c.zl, c.zr)

		cubes := []Cube{}
		for i := 0; i < len(xsplit); i += 2 {
			for j := 0; j < len(ysplit); j += 2 {
				for k := 0; k < len(zsplit); k += 2 {
					cubes = append(cubes, Cube{xsplit[i], xsplit[i+1], ysplit[j], ysplit[j+1], zsplit[k], zsplit[k+1]})
				}
			}
		}

		for _, x := range cubes {
			if x == c {
				continue
			}
			ret = append(ret, x)
		}
		return
	}

	volumn := func(c Cube) int {
		return (c.xr - c.xl + 1) * (c.yr - c.yl + 1) * (c.zr - c.zl + 1)
	}

	removeElements := func(arr []Cube, indexes []int) (ret []Cube) {
		hit := map[int]bool{}
		for _, i := range indexes {
			hit[i] = true
		}

		for i, e := range arr {
			if !hit[i] {
				ret = append(ret, e)
			}
		}
		return
	}

	cubes := []Cube{}
	for _, l := range lines {
		op := strings.Fields(l)[0]

		r, _ := regexp.Compile("x=(.*?),y=(.*?),z=(.*)")
		f := r.FindAllStringSubmatch(l, -1)

		xRange, yRange, zRange := f[0][1], f[0][2], f[0][3]

		xl, xr := getRange(xRange)
		yl, yr := getRange(yRange)
		zl, zr := getRange(zRange)

		next := []Cube{}

		cube := Cube{xl, xr, yl, yr, zl, zr}
		if op == "on" {
			next = append(next, cubes...)
			next = append(next, cube)
		} else {
			for _, c := range cubes {
				next = append(next, subtractCubes(c, cube)...)
			}
		}
		cubes = next
	}

	startPtr := 0
	for {
		overlappedCubesMap := map[Cube]bool{}
		endPtr := startPtr
		for i := startPtr; i < len(cubes); i++ {
			for j := i + 1; j < len(cubes); j++ {
				c, o := overlappedCube(cubes[i], cubes[j])
				if o {
					c1 := subtractCubes(cubes[i], c)
					c2 := subtractCubes(cubes[j], c)
					cubes = removeElements(cubes, []int{i, j})
					cubes = append(cubes, c1...)
					cubes = append(cubes, c2...)
					overlappedCubesMap[c] = true
				} else if len(overlappedCubesMap) == 0 {
					endPtr = i
				}
			}
		}
		if len(overlappedCubesMap) <= 0 {
			break
		}

		arr := []Cube{}
		for k := range overlappedCubesMap {
			arr = append(arr, k)
		}

		cubes = append(cubes, arr...)
		startPtr = endPtr

	}

	v := 0
	for _, c := range cubes {
		v += volumn(c)
	}

	return v
}

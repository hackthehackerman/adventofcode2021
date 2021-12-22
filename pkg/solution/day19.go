package solution

import (
	"math"
	"strconv"
	"strings"
)

type Coordinate3D struct {
	x, y, z int
}

func rotate(start Coordinate3D) []Coordinate3D {
	x, y, z := start.x, start.y, start.z
	return []Coordinate3D{
		{x, y, z},
		{y, -x, z},
		{-x, -y, z},
		{-y, x, z},
		{y, z, x},
		{-x, z, y},
		{-y, z, -x},
		{x, z, -y},
		{z, x, y},
		{z, y, -x},
		{z, -x, -y},
		{z, -y, x},
		{y, x, -z},
		{-x, y, -z},
		{-y, -x, -z},
		{x, -y, -z},
		{x, -z, y},
		{y, -z, -x},
		{-x, -z, -y},
		{-y, -z, x},
		{-z, y, x},
		{-z, -x, y},
		{-z, -y, -x},
		{-z, x, -y},
	}
}

func rotateBeacons(b []Coordinate3D) [][]Coordinate3D {
	tmp := [][]Coordinate3D{}
	for _, c := range b {
		tmp = append(tmp, rotate(c))
	}

	ret := [][]Coordinate3D{}
	for i := 0; i < 24; i++ {
		nb := []Coordinate3D{}
		for _, c := range tmp {
			nb = append(nb, c[i])
		}
		ret = append(ret, nb)
	}
	return ret
}

func addCoordinates(c1, c2 Coordinate3D) (ret Coordinate3D) {
	ret.x = c1.x + c2.x
	ret.y = c1.y + c2.y
	ret.z = c1.z + c2.z
	return
}

func minusCoordinates(c1, c2 Coordinate3D) (ret Coordinate3D) {
	ret.x = c1.x - c2.x
	ret.y = c1.y - c2.y
	ret.z = c1.z - c2.z
	return
}

func relative(b1, b2 []Coordinate3D) (rc Coordinate3D, rotation int, found bool) {
	b1g := rotateBeacons(b1)
	for r, b1Rotated := range b1g {
		for _, b1c := range b1Rotated {
			for _, b2c := range b2 {
				distance := minusCoordinates(b2c, b1c)
				translated := make(map[Coordinate3D]bool)
				for _, b1cc := range b1Rotated {
					t := addCoordinates(b1cc, distance)
					translated[t] = true
				}
				matchCnt := 0
				for _, b2cc := range b2 {
					if translated[b2cc] {
						matchCnt += 1
					}
				}
				if matchCnt >= 12 {
					// found!
					return distance, r, true
				}
			}
		}
	}
	return Coordinate3D{}, 0, false
}

func (s *Solution) Day19Part1(fn string) (ret int) {
	lines := toLines(fn)

	// 1. parse all beacons coordinates
	beacons := [][]Coordinate3D{}
	buffer := []Coordinate3D{}
	for i, l := range lines {
		split := strings.Split(l, ",")
		if len(split) == 3 {
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			z, _ := strconv.Atoi(split[2])
			buffer = append(buffer, Coordinate3D{x, y, z})
		}
		if l == "" || i == len(lines)-1 {
			beacons = append(beacons, buffer)
			buffer = []Coordinate3D{}
		}
	}

	// 2. For each scanner, found whether they overlapped with each other, and the scanner coordinate differences
	// between them
	type Difference struct {
		translation Coordinate3D
		rotation    int
	}

	paths := make(map[Point]Difference)
	for i := 0; i < len(beacons); i++ {
		for j := 0; j < len(beacons); j++ {
			if i == j {
				continue
			}

			if rc, rotation, found := relative(beacons[i], beacons[j]); found == true {
				paths[Point{i, j}] = Difference{rc, rotation}
			}
		}
	}

	// 3. For each scanner, try to normalize their beacons coordinates into becon0's sytem
	var dfs func(start, end, max int, m map[Point]Difference, visited map[int]bool) []int
	dfs = func(start, end, max int, m map[Point]Difference, visited map[int]bool) []int {
		if start == end {
			return []int{start}
		}
		if _, found := m[Point{start, end}]; found {
			return []int{start, end}
		}
		visited[start] = true
		for i := 0; i < max; i++ {
			if i == start || visited[i] {
				continue
			}
			if _, found := m[Point{start, i}]; found {

				if rest := dfs(i, end, max, m, visited); len(rest) > 0 {
					visited[start] = false
					return append([]int{start}, rest...)
				}
			}
		}
		return []int{}
	}

	translate := func(b []Coordinate3D, d Difference) []Coordinate3D {
		rotated := rotateBeacons(b)[d.rotation]
		ret := []Coordinate3D{}
		for _, r := range rotated {
			translated := addCoordinates(r, d.translation)
			ret = append(ret, translated)
		}
		return ret
	}

	normalizedBeacons := make(map[Coordinate3D]bool)
	for i := 0; i < len(beacons); i++ {
		b := make([]Coordinate3D, len(beacons[i]))
		copy(b, beacons[i])
		p := dfs(i, 0, len(beacons), paths, make(map[int]bool))

		for j := 0; j < len(p)-1; j++ {
			d := paths[Point{p[j], p[j+1]}]
			b = translate(b, d)
		}

		for _, bb := range b {
			normalizedBeacons[bb] = true
		}
	}

	return len(normalizedBeacons)
}

func (s *Solution) Day19Part2(fn string) (ret int) {
	lines := toLines(fn)

	// 1. parse all beacons coordinates
	beacons := [][]Coordinate3D{}
	buffer := []Coordinate3D{}
	for i, l := range lines {
		split := strings.Split(l, ",")
		if len(split) == 3 {
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			z, _ := strconv.Atoi(split[2])
			buffer = append(buffer, Coordinate3D{x, y, z})
		}
		if l == "" || i == len(lines)-1 {
			beacons = append(beacons, buffer)
			buffer = []Coordinate3D{}
		}
	}

	// 2. For each scanner, found whether they overlapped with each other, and the scanner coordinate differences
	// between them
	type Difference struct {
		translation Coordinate3D
		rotation    int
	}

	paths := make(map[Point]Difference)
	for i := 0; i < len(beacons); i++ {
		for j := 0; j < len(beacons); j++ {
			if i == j {
				continue
			}

			if rc, rotation, found := relative(beacons[i], beacons[j]); found == true {
				paths[Point{i, j}] = Difference{rc, rotation}
			}
		}
	}

	// 3. For each scanner, try to normalize their beacons coordinates into becon0's sytem
	var dfs func(start, end, max int, m map[Point]Difference, visited map[int]bool) []int
	dfs = func(start, end, max int, m map[Point]Difference, visited map[int]bool) []int {
		if start == end {
			return []int{start}
		}
		if _, found := m[Point{start, end}]; found {
			return []int{start, end}
		}
		visited[start] = true
		for i := 0; i < max; i++ {
			if i == start || visited[i] {
				continue
			}
			if _, found := m[Point{start, i}]; found {

				if rest := dfs(i, end, max, m, visited); len(rest) > 0 {
					visited[start] = false
					return append([]int{start}, rest...)
				}
			}
		}
		return []int{}
	}

	beaconCoordinates := []Coordinate3D{{}}
	for i := 1; i < len(beacons); i++ {
		c := Coordinate3D{}
		p := dfs(i, 0, len(beacons), paths, make(map[int]bool))

		for j := 0; j < len(p)-1; j++ {
			d := paths[Point{p[j], p[j+1]}]
			c = rotate(c)[d.rotation]
			c = addCoordinates(c, d.translation)
		}

		beaconCoordinates = append(beaconCoordinates, c)
	}

	maxDistance := 0
	for i := 0; i < len(beaconCoordinates)-1; i++ {
		for j := i + 1; j < len(beaconCoordinates); j++ {
			bc1, bc2 := beaconCoordinates[i], beaconCoordinates[j]
			d := math.Abs(float64(bc1.x-bc2.x)) + math.Abs(float64(bc1.y-bc2.y)) + math.Abs(float64(bc1.z-bc2.z))
			if int(d) > maxDistance {
				maxDistance = int(d)
			}
		}
	}

	return maxDistance
}

package solution

import (
	"strconv"
	"strings"
)

func enhance(image [][]int, cipher []int, defaultVal int) (ret [][]int, newDefaultVal int) {
	for i := -1; i < len(image)+1; i++ {
		row := []int{}
		for j := -1; j < len(image[0])+1; j++ {
			pos := []string{}
			for ii := i - 1; ii < i+2; ii++ {
				for jj := j - 1; jj < j+2; jj++ {
					if ii < 0 || ii >= len(image) || jj < 0 || jj >= len(image[0]) {
						pos = append(pos, strconv.Itoa(defaultVal))
					} else {
						pos = append(pos, strconv.Itoa(image[ii][jj]))
					}
				}
			}
			posInt, _ := strconv.ParseInt(strings.Join(pos, ""), 2, 64)
			row = append(row, cipher[posInt])
		}
		ret = append(ret, row)
	}

	newDefaultValBinary := []string{}
	for i := 0; i < 9; i++ {
		newDefaultValBinary = append(newDefaultValBinary, strconv.Itoa(defaultVal))
	}
	defaultValPosInt, _ := strconv.ParseInt(strings.Join(newDefaultValBinary, ""), 2, 64)
	newDefaultVal = cipher[defaultValPosInt]

	return ret, newDefaultVal
}

func (s *Solution) Day20Part1(fn string) (ret int) {
	lines := toLines(fn)

	cipher := []int{}
	image := [][]int{}

	writeImage := false
	for _, l := range lines {
		if l == "" {
			writeImage = true
			continue
		}

		row := []int{}
		for _, s := range l {
			switch s {
			case '.':
				row = append(row, 0)
			case '#':
				row = append(row, 1)
			}
		}

		if writeImage {
			image = append(image, row)
		} else {
			cipher = append(cipher, row...)
		}
	}

	post1, defaultVal := enhance(image, cipher, 0)
	post2, _ := enhance(post1, cipher, defaultVal)

	cnt := 0
	for _, row := range post2 {
		for _, n := range row {
			cnt += n
		}
	}

	return cnt
}

func (s *Solution) Day20Part2(fn string) (ret int) {
	lines := toLines(fn)

	cipher := []int{}
	image := [][]int{}

	writeImage := false
	for _, l := range lines {
		if l == "" {
			writeImage = true
			continue
		}

		row := []int{}
		for _, s := range l {
			switch s {
			case '.':
				row = append(row, 0)
			case '#':
				row = append(row, 1)
			}
		}

		if writeImage {
			image = append(image, row)
		} else {
			cipher = append(cipher, row...)
		}
	}

	post := image
	defaultVal := 0
	for i := 0; i < 50; i++ {
		post, defaultVal = enhance(post, cipher, defaultVal)
	}

	cnt := 0
	for _, row := range post {
		for _, n := range row {
			cnt += n
		}
	}

	return cnt
}

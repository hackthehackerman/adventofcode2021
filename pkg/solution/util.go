package solution

import (
	"bufio"
	"os"
	"strconv"
)

func toLines(fn string) (r []string) {
	file, _ := os.Open(fn)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	return
}

func toInts(l []string) (r []int) {
	for _, s := range l {
		n, _ := strconv.Atoi(s)
		r = append(r, n)
	}
	return r
}

func sort(a int, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func max(ints []int) int {
	if len(ints) <= 0 {
		return 0
	}
	m := ints[0]
	for _, v := range ints {
		if v > m {
			m = v
		}
	}
	return m
}

func min(ints []int) int {
	if len(ints) <= 0 {
		return 0
	}
	m := ints[0]
	for _, v := range ints {
		if v < m {
			m = v
		}
	}
	return m
}

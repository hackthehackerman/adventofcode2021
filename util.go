package main

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

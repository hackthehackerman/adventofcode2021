package main

import (
	"fmt"
	"strconv"
)

func dayOne(p string) {
	switch p {
	case "1":
		d1p1()
	case "2":
		d1p2()
	}
}

func d1p1() {
	lines := toLines("f1_1")
	prev := 0
	cnt := 0
	for i, s := range lines {
		n, _ := strconv.Atoi(s)
		if i > 0 && n > prev {
			cnt += 1
		}
		prev = n
	}
	fmt.Println(cnt)
}

func d1p2() {
	nums := toInts(toLines("f1_1"))
	cnt := 0
	for i, n := range nums {
		if i > 2 && n > nums[i-3] {
			cnt += 1
		}
	}
	fmt.Println(cnt)
}

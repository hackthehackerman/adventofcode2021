package solution

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
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

func sortInts(a int, b int) (int, int) {
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

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func Perm(a []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(a, len(a))
	return res
}

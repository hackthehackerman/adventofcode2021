package solution

import (
	"fmt"
	"math"
	"strconv"
)

type FishPair struct {
	lFish *FishPair
	rFish *FishPair
	lInt  int
	rInt  int
	depth int
}

func parse(s string) (fp FishPair) {
	stack := []*FishPair{}
	curInt := 0
	for i, r := range s {
		switch r {
		case '[':
			depth := 0
			if len(stack) != 0 {
				depth = stack[len(stack)-1].depth + 1
			}
			newFp := FishPair{depth: depth}
			stack = append(stack, &newFp)
		case ']':
			if s[i-1:i] == "]" {
				newFp := stack[len(stack)-1]
				stack[len(stack)-2].rFish = newFp
				stack = stack[:len(stack)-1]
			} else {
				stack[len(stack)-1].rInt = curInt
				curInt = 0
			}
		case ',':
			if s[i-1:i] == "]" {
				newFp := stack[len(stack)-1]
				stack[len(stack)-2].lFish = newFp
				stack = stack[:len(stack)-1]
			} else {
				stack[len(stack)-1].lInt = curInt
				curInt = 0
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n, _ := strconv.Atoi(string(r))
			curInt = curInt*10 + n
		default:
			panic(fmt.Sprintf("%s not supported", string(r)))
		}
	}
	return *stack[0]
}

func add(f1, f2 *FishPair) (fpp *FishPair) {
	// increment all depth by 1
	var walk func(f *FishPair)
	walk = func(f *FishPair) {
		if f.lFish != nil {
			walk(f.lFish)
		}
		if f.rFish != nil {
			walk(f.rFish)
		}
		f.depth += 1
	}

	walk(f1)
	walk(f2)

	fp := &FishPair{}
	fp.lFish = f1
	fp.rFish = f2
	return fp
}

func explode(f *FishPair) bool {
	flatten := []*FishPair{}
	var walk func(f *FishPair)
	walk = func(f *FishPair) {
		if f.lFish != nil && f.rFish != nil {
			walk(f.lFish)
			walk(f.rFish)
		} else if f.lFish != nil {
			walk(f.lFish)
			flatten = append(flatten, f)
		} else if f.rFish != nil {
			flatten = append(flatten, f)
			walk(f.rFish)
		} else {
			flatten = append(flatten, f)
		}
	}
	walk(f)

	var findParent func(f, child *FishPair) *FishPair
	findParent = func(f, child *FishPair) *FishPair {
		if f.lFish != nil && f.lFish == child {
			return f
		}
		if f.rFish != nil && f.rFish == child {
			return f
		}
		if f.lFish != nil {
			if findParent(f.lFish, child) != nil {
				return findParent(f.lFish, child)
			}
		}
		if f.rFish != nil {
			if findParent(f.rFish, child) != nil {
				return findParent(f.rFish, child)
			}
		}
		return nil
	}

	exploded := false
	for i, fp := range flatten {
		if exploded {
			break
		}

		if fp.depth >= 4 && fp.lFish == nil && fp.rFish == nil {
			if i > 0 {
				if flatten[i-1].rFish == nil {
					flatten[i-1].rInt += fp.lInt
				} else {
					flatten[i-1].lInt += fp.lInt
				}
			}
			if i < len(flatten)-1 {
				if flatten[i+1].lFish == nil {
					flatten[i+1].lInt += fp.rInt
				} else {
					flatten[i+1].rInt += fp.rInt
				}
			}

			parent := findParent(f, fp)
			if parent.lFish == fp {
				parent.lFish = nil
			} else if parent.rFish == fp {
				parent.rFish = nil
			}
			exploded = true
		}
	}
	return exploded
}

func split(f *FishPair) bool {
	flatten := []*FishPair{}
	var walk func(f *FishPair)
	walk = func(f *FishPair) {
		if f.lFish != nil && f.rFish != nil {
			walk(f.lFish)
			walk(f.rFish)
		} else if f.lFish != nil {
			walk(f.lFish)
			flatten = append(flatten, f)
		} else if f.rFish != nil {
			flatten = append(flatten, f)
			walk(f.rFish)
		} else {
			flatten = append(flatten, f)
		}
	}
	walk(f)
	splited := false
	for _, fp := range flatten {
		if splited {
			break
		}

		if fp.lFish == nil && fp.lInt >= 10 {
			// split
			lfp := FishPair{depth: fp.depth + 1}
			lfp.lInt = fp.lInt / 2
			lfp.rInt = (fp.lInt + 1) / 2
			// lfp.parent = fp
			fp.lInt = 0
			fp.lFish = &lfp
			splited = true
		} else if fp.rFish == nil && fp.rInt >= 10 {
			// split
			rfp := FishPair{depth: fp.depth + 1}
			rfp.lInt = fp.rInt / 2
			rfp.rInt = (fp.rInt + 1) / 2
			// rfp.parent = fp
			fp.rInt = 0
			fp.rFish = &rfp
			splited = true
		}
	}
	return splited
}

func reduce(f *FishPair) {
	var walk func(f *FishPair)
	walk = func(f *FishPair) {

		if f.lFish != nil && f.rFish != nil {
			walk(f.lFish)
			walk(f.rFish)
		} else if f.lFish != nil {
			walk(f.lFish)
		} else if f.rFish != nil {
			walk(f.rFish)
		}
	}
	walk(f)

	for {
		reduced := explode(f)
		if reduced {
			continue
		}
		reduced = split(f)

		if !reduced {
			break
		}
	}
}

func printFish(fp *FishPair) {
	// return
	prefix := "  "
	for i := 0; i < fp.depth; i++ {
		prefix += "  "
	}
	fmt.Println(prefix, &fp, fp, fp.lFish, fp.rFish)
	if fp.lFish != nil {
		printFish(fp.lFish)
	}
	if fp.rFish != nil {
		printFish(fp.rFish)
	}
}

func printFishLine(fp *FishPair) {
	print("[")
	if fp.lFish != nil {
		printFishLine(fp.lFish)
	} else {
		print(fp.lInt)
	}
	print(",")
	if fp.rFish != nil {
		printFishLine(fp.rFish)
	} else {
		print(fp.rInt)
	}
	print("]")
}

func magnitude(fp *FishPair) int {
	lv := fp.lInt
	rv := fp.rInt
	if fp.lFish != nil {
		lv = magnitude(fp.lFish)
	}
	if fp.rFish != nil {
		rv = magnitude(fp.rFish)
	}
	return lv*3 + rv*2
}

func (s *Solution) Day18Part1(fn string) (ret int) {
	lines := toLines(fn)

	var fp *FishPair
	for _, l := range lines {
		nFP := parse(l)
		if fp == nil {
			fp = &nFP
		} else {
			fp = add(fp, &nFP)
		}
		reduce(fp)
	}

	return magnitude(fp)
}

func (s *Solution) Day18Part2(fn string) (ret int) {
	lines := toLines(fn)

	maxMagnitude := math.MinInt
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}

			fpl := parse(lines[i])
			fpr := parse(lines[j])

			added := add(&fpl, &fpr)
			reduce(added)
			mag := magnitude(added)
			if mag > maxMagnitude {
				maxMagnitude = mag
			}
		}
	}

	return maxMagnitude
}

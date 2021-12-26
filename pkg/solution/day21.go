package solution

import (
	"strconv"
	"strings"
)

func (s *Solution) Day21Part1(fn string) (ret int) {
	lines := toLines(fn)

	split0 := strings.Fields(lines[0])
	split1 := strings.Fields(lines[1])

	start0, _ := strconv.Atoi(split0[len(split0)-1])
	start1, _ := strconv.Atoi(split1[len(split1)-1])

	score0 := 0
	score1 := 0
	numRolls := 1
	for {
		numRolls += 3
		if map[int]bool{4: true, 5: true, 0: true}[numRolls%6] {
			// player 0
			start0 = (start0 + (numRolls-3)*3 + 3)
			for start0 > 10 {
				start0 = start0 - 10
			}
			score0 += start0
			if score0 >= 1000 {
				break
			}
		} else {
			// player 1
			start1 = (start1 + (numRolls-3)*3 + 3)
			for start1 > 10 {
				start1 = start1 - 10
			}
			score1 += start1
			if score1 >= 1000 {
				break
			}
		}
	}

	if score0 >= 1000 {
		return score1 * (numRolls - 1)
	} else {
		return score0 * (numRolls - 1)
	}
}

func (s *Solution) Day21Part2(fn string) (ret int) {
	lines := toLines(fn)

	split0 := strings.Fields(lines[0])
	split1 := strings.Fields(lines[1])

	start0, _ := strconv.Atoi(split0[len(split0)-1])
	start1, _ := strconv.Atoi(split1[len(split1)-1])

	type Board struct {
		pos0, score0, pos1, score1 int
	}

	universe := map[Board]int{
		{start0, 0, start1, 0}: 1,
	}
	win0 := 0
	win1 := 0
	for i := 1; ; i++ {
		next := map[Board]int{}
		// notDone := false
		for board, count := range universe {
			for i := 1; i <= 3; i++ {
				for j := 1; j <= 3; j++ {
					for k := 1; k <= 3; k++ {
						nb := board
						nb.pos0 = nb.pos0 + i + j + k

						for nb.pos0 > 10 {
							nb.pos0 -= 10
						}
						nb.score0 += nb.pos0

						if nb.score0 >= 21 {
							win0 += count
							continue
						}

						for ii := 1; ii <= 3; ii++ {
							for jj := 1; jj <= 3; jj++ {
								for kk := 1; kk <= 3; kk++ {
									nnb := nb

									nnb.pos1 = nb.pos1 + ii + jj + kk
									for nnb.pos1 > 10 {
										nnb.pos1 -= 10
									}

									nnb.score1 += nnb.pos1

									if nnb.score1 >= 21 {
										win1 += count
									} else {
										next[nnb] += count
									}
								}
							}
						}
					}
				}
			}
		}
		universe = next

		if len(universe) == 0 {
			break
		}
	}

	if win0 > win1 {
		return win0
	} else {
		return win1
	}
}

package solution

import (
	"fmt"
	"math"
)

type FishBoard struct {
	hallway [11]string
	rooms   [4][4]string
}

func solve(startingBoard FishBoard, roomLength int) (ret int) {
	type State struct {
		board FishBoard
		cost  int
	}

	roomIdx := func(f string) int {
		switch f {
		case "A":
			return 0
		case "B":
			return 1
		case "C":
			return 2
		case "D":
			return 3
		}
		panic("Unsuppored fish")
		return -1
	}

	moveCost := func(f string) int {
		switch f {
		case "A":
			return 1
		case "B":
			return 10
		case "C":
			return 100
		case "D":
			return 1000
		}
		panic("Unsuppored fish")
		return -1
	}

	legal := func(b FishBoard) bool {
		cnt := map[string]int{}
		for _, v := range b.hallway {
			cnt[v] += 1
		}

		for i := 0; i < 4; i++ {
			for j := 0; j < roomLength; j++ {
				cnt[b.rooms[i][j]] += 1
			}
		}

		return cnt["A"] == roomLength && cnt["B"] == roomLength && cnt["C"] == roomLength && cnt["D"] == roomLength
	}

	minCostToEnd := func(b FishBoard) (cost int) {
		for i := 0; i < 11; i++ {
			f := b.hallway[i]
			if f == "" {
				continue
			}
			ri := roomIdx(f)

			hallwayEntranceIdx := 2 + (ri * 2)
			s := min([]int{i, hallwayEntranceIdx})
			e := max([]int{i, hallwayEntranceIdx})

			for roomIdx := roomLength - 1; roomIdx >= 0; roomIdx-- {
				if b.rooms[ri][roomIdx] != f {
					b.rooms[ri][roomIdx] = f
					cost += ((e - s) + (roomIdx + 1)) * moveCost(f)
					break
				}
			}
		}

		for i := 0; i < 4; i++ {
			for j := 0; j < roomLength; j++ {
				f := b.rooms[i][j]
				if f == "" {
					continue
				}
				ri := roomIdx(f)

				if ri == i {
					continue
				}

				s := min([]int{i, ri})
				e := max([]int{i, ri})

				hallwayDistance := 2 * (e - s)

				for roomIdx := roomLength - 1; roomIdx >= 0; roomIdx-- {
					if b.rooms[ri][roomIdx] != f {
						b.rooms[ri][roomIdx] = f
						cost += ((j + 1) + hallwayDistance + roomIdx + 1) * moveCost(f)
						break
					}
				}
			}
		}
		return cost
	}

	canMoveToRoom := func(hIdx, roomPos int, b FishBoard) (canMove bool, cost int) {
		f := b.hallway[hIdx]
		ri := roomIdx(f)

		hallwayEntranceIdx := 2 + (ri * 2)
		s := min([]int{hIdx, hallwayEntranceIdx})
		e := max([]int{hIdx, hallwayEntranceIdx})

		for i := s + 1; i < e; i++ {
			if b.hallway[i] != "" {
				return false, 0
			}
		}

		for r := roomPos; r >= 0; r-- {
			if b.rooms[ri][r] != "" {
				return false, 0
			}
		}

		for r := roomPos + 1; r < roomLength; r++ {
			if b.rooms[ri][r] != f {
				return false, 0
			}
		}

		return true, ((e - s) + (roomPos + 1)) * moveCost(f)
	}

	moveToRoom := func(hIdx, roomPos int, b FishBoard) (ret FishBoard) {
		ret = b
		f := ret.hallway[hIdx]
		ri := roomIdx(f)

		ret.rooms[ri][roomPos] = f
		ret.hallway[hIdx] = ""
		return
	}

	canMoveToHallway := func(rIdx, roomPos, hIdx int, b FishBoard) (canMove bool, cost int) {
		f := b.rooms[rIdx][roomPos]
		ri := roomIdx(f)

		hallwayEntranceIdx := 2 + (rIdx * 2)
		s := min([]int{hIdx, hallwayEntranceIdx})
		e := max([]int{hIdx, hallwayEntranceIdx})
		for i := s + 1; i < e; i++ {
			if b.hallway[i] != "" {
				return false, 0
			}
		}

		if hIdx == 2 || hIdx == 4 || hIdx == 6 || hIdx == 8 {
			return false, 0
		}

		if b.hallway[hIdx] != "" {
			return false, 0
		}

		for r := roomPos - 1; r >= 0; r-- {
			if b.rooms[rIdx][r] != "" {
				return false, 0
			}
		}

		if ri == rIdx {
			filled := true
			for r := roomPos + 1; r < roomLength; r++ {
				if b.rooms[ri][r] != f {
					filled = false
				}
			}
			if filled {
				return false, 0
			}
		}

		return true, ((e - s) + (roomPos + 1)) * moveCost(f)
	}

	moveToHallway := func(rIdx, roomPos, hIdx int, b FishBoard) (ret FishBoard) {
		ret = b
		f := ret.rooms[rIdx][roomPos]

		ret.hallway[hIdx] = f
		ret.rooms[rIdx][roomPos] = ""

		if !legal(ret) {
			fmt.Println(rIdx, roomPos, hIdx, b, ret)
			panic("haha")
		}
		return
	}

	done := func(b FishBoard) bool {
		expected := [4]string{"A", "B", "C", "D"}
		for i := 0; i < 4; i++ {
			for j := 0; j < roomLength; j++ {
				if b.rooms[i][j] != expected[i] {
					return false
				}
			}
		}
		return true
	}

	states := []State{
		{startingBoard, 0},
	}

	minCost := math.MaxInt
	for {
		if len(states) == 0 {
			break
		}

		curState := states[len(states)-1]
		states = states[:len(states)-1]

		if done(curState.board) {
			if curState.cost < minCost {
				minCost = curState.cost
			}
			continue
		}

		// move fish from hallway to rooms
		for i := 0; i < 11; i++ {
			f := curState.board.hallway[i]
			if f == "" {
				continue
			}

			for roomPos := 0; roomPos < roomLength; roomPos++ {
				canMove, cost := canMoveToRoom(i, roomPos, curState.board)

				if canMove {

					nb := moveToRoom(i, roomPos, curState.board)
					if curState.cost+cost+minCostToEnd(nb) < minCost {
						states = append(states, State{nb, curState.cost + cost})
					}
				}
			}
		}

		// move fish from rooms to hallway
		for i := 0; i < 4; i++ {
			for roomPos := 0; roomPos < roomLength; roomPos++ {
				f := curState.board.rooms[i][roomPos]
				if f == "" {
					continue
				}

				for hi := 0; hi < 11; hi++ {
					canMove, cost := canMoveToHallway(i, roomPos, hi, curState.board)
					if canMove {
						nb := moveToHallway(i, roomPos, hi, curState.board)
						if curState.cost+cost+minCostToEnd(nb) < minCost {
							states = append(states, State{nb, curState.cost + cost})
						}
					}
				}
			}
		}
	}

	return minCost
}

func (s *Solution) Day23Part1(fn string) (ret int) {
	lines := toLines(fn)
	r1, r2, r3, r4 := lines[2][3:4], lines[2][5:6], lines[2][7:8], lines[2][9:10]
	r5, r6, r7, r8 := lines[3][3:4], lines[3][5:6], lines[3][7:8], lines[3][9:10]

	startingBoard := FishBoard{
		rooms: [4][4]string{{r1, r5, "A", "A"}, {r2, r6, "B", "B"}, {r3, r7, "C", "C"}, {r4, r8, "D", "D"}},
	}

	return solve(startingBoard, 2)
}

func (s *Solution) Day23Part2(fn string) (ret int) {
	lines := toLines(fn)
	r1, r2, r3, r4 := lines[2][3:4], lines[2][5:6], lines[2][7:8], lines[2][9:10]
	r5, r6, r7, r8 := "D", "C", "B", "A"
	r9, r10, r11, r12 := "D", "B", "A", "C"
	r13, r14, r15, r16 := lines[3][3:4], lines[3][5:6], lines[3][7:8], lines[3][9:10]

	startingBoard := FishBoard{
		rooms: [4][4]string{{r1, r5, r9, r13}, {r2, r6, r10, r14}, {r3, r7, r11, r15}, {r4, r8, r12, r16}},
	}

	return solve(startingBoard, 4)
}

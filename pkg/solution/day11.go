package solution

func (s *Solution) Day11Part1(fn string) (ret int) {
	increment := func(m [][]int, i, j, numRow, numColumn int) {
		if i < 0 ||
			i >= numRow ||
			j < 0 ||
			j >= numColumn {
			return
		}
		m[i][j] += 1
	}

	m := toMatrix(fn)
	numRow := len(m)
	numColumn := len(m[0])
	cnt := 0
	for step := 0; step < 100; step++ {
		flashed := make(map[Point]bool)

		// add 1 to all jelly fish
		for i := 0; i < numRow; i++ {
			for j := 0; j < numRow; j++ {
				m[i][j] += 1
			}
		}

		// flash
		numFlashed := len(flashed)
		for {
			for i := 0; i < numRow; i++ {
				for j := 0; j < numRow; j++ {
					if m[i][j] > 9 && !flashed[Point{i, j}] {
						cnt += 1
						flashed[Point{i, j}] = true
						increment(m, i-1, j-1, numRow, numColumn) // tl
						increment(m, i-1, j, numRow, numColumn)   // top
						increment(m, i+1, j-1, numRow, numColumn) // tr
						increment(m, i, j+1, numRow, numColumn)   // right
						increment(m, i+1, j+1, numRow, numColumn) // br
						increment(m, i+1, j, numRow, numColumn)   // bottom
						increment(m, i-1, j+1, numRow, numColumn) // bl
						increment(m, i, j-1, numRow, numColumn)   // left
					}
				}
			}
			if len(flashed) == numFlashed {
				break
			}
			numFlashed = len(flashed)
		}
		// zero
		for i := 0; i < numRow; i++ {
			for j := 0; j < numRow; j++ {
				if m[i][j] > 9 {
					m[i][j] = 0
				}
			}
		}
	}
	return cnt
}

func (s *Solution) Day11Part2(fn string) (ret int) {
	increment := func(m [][]int, i, j, numRow, numColumn int) {
		if i < 0 ||
			i >= numRow ||
			j < 0 ||
			j >= numColumn {
			return
		}
		m[i][j] += 1
	}

	m := toMatrix(fn)
	numRow := len(m)
	numColumn := len(m[0])
	for step := 0; ; step++ {
		flashed := make(map[Point]bool)

		// add 1 to all jelly fish
		for i := 0; i < numRow; i++ {
			for j := 0; j < numRow; j++ {
				m[i][j] += 1
			}
		}

		// flash
		numFlashed := len(flashed)
		for {
			for i := 0; i < numRow; i++ {
				for j := 0; j < numRow; j++ {
					if m[i][j] > 9 && !flashed[Point{i, j}] {
						flashed[Point{i, j}] = true
						increment(m, i-1, j-1, numRow, numColumn) // tl
						increment(m, i-1, j, numRow, numColumn)   // top
						increment(m, i+1, j-1, numRow, numColumn) // tr
						increment(m, i, j+1, numRow, numColumn)   // right
						increment(m, i+1, j+1, numRow, numColumn) // br
						increment(m, i+1, j, numRow, numColumn)   // bottom
						increment(m, i-1, j+1, numRow, numColumn) // bl
						increment(m, i, j-1, numRow, numColumn)   // left
					}
				}
			}
			if len(flashed) == numFlashed {
				break
			}
			numFlashed = len(flashed)
		}
		if len(flashed) == numRow*numColumn {
			return step + 1
		}

		// zero
		for i := 0; i < numRow; i++ {
			for j := 0; j < numRow; j++ {
				if m[i][j] > 9 {
					m[i][j] = 0
				}
			}
		}
	}
}

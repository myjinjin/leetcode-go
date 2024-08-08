package leetcode

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	result := [][]int{{rStart, cStart}}
	if rows*cols == 1 {
		return result
	}

	steps := 0
	r, c := rStart, cStart

	for len(result) < rows*cols {
		for d := 0; d < 4; d++ {
			if d%2 == 0 {
				steps++
			}
			for i := 0; i < steps; i++ {
				r += directions[d][0]
				c += directions[d][1]

				if r >= 0 && r < rows && c >= 0 && c < cols {
					result = append(result, []int{r, c})
				}
				if len(result) == rows*cols {
					return result
				}
			}
		}
	}

	return result
}

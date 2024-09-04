package leetcode

type Point struct {
	x int
	y int
}

func robotSim(commands []int, obstacles [][]int) int {
	far := 0

	obstacleSet := make(map[Point]bool)
	for _, o := range obstacles {
		x, y := o[0], o[1]
		p := Point{x: x, y: y}
		obstacleSet[p] = true
	}

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	dirIdx := 0
	currX, currY := 0, 0

	for _, c := range commands {
		if c == -2 {
			dirIdx = (dirIdx + 3) % 4
		} else if c == -1 {
			dirIdx = (dirIdx + 1) % 4
		} else {
			d := directions[dirIdx]
			x, y := d[0], d[1]
			for i := 0; i < c; i++ {
				if _, exist := obstacleSet[Point{currX + x, currY + y}]; exist {
					break
				}
				currX += x
				currY += y
				far = max(far, currX*currX+currY*currY)
			}
		}
	}

	return far
}

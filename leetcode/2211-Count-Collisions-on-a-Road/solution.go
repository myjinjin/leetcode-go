package leetcode

func countCollisions(directions string) int {
	collisions := 0
	right := 0
	hasStaying := false

	for i := 0; i < len(directions); i++ {
		switch directions[i] {
		case 'R':
			right++
		case 'S':
			collisions += right
			right = 0
			hasStaying = true
		case 'L':
			if right > 0 {
				collisions += right + 1
				right = 0
				hasStaying = true
			} else if hasStaying {
				collisions++
			}
		}
	}

	return collisions
}

package leetcode

func predictPartyVictory(senate string) string {
	n := len(senate)
	radiant := make([]int, 0)
	dire := make([]int, 0)

	for i := 0; i < n; i++ {
		if senate[i] == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+n)
		} else {
			dire = append(dire, dire[0]+n)
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}

	if len(radiant) > 0 {
		return "Radiant"
	}
	return "Dire"
}

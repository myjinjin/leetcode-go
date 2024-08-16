package leetcode

func addMinimum(word string) int {
	j := 0
	add := 0
	for i := 0; i < len(word); i++ {
		for j%3 != int(word[i]-'a') {
			add++
			j++
		}
		j++
	}

	if j%3 != 0 {
		add += (3 - j%3)
	}
	return add
}

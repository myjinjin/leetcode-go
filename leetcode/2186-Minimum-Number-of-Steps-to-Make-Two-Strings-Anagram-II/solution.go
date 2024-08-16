package leetcode

func minSteps(s string, t string) int {
	steps := 0

	counts := [26]int{}
	for i := 0; i < len(s); i++ {
		counts[int(s[i]-'a')]++
	}
	for i := 0; i < len(t); i++ {
		counts[int(t[i]-'a')]--
	}

	for i := 0; i < 26; i++ {
		if counts[i] != 0 {
			steps += abs(counts[i])
		}
	}

	return steps
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

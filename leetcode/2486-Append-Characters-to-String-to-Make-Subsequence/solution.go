package leetcode

func appendCharacters(s string, t string) int {
	tIdx := 0
	sIdx := 0
	for tIdx < len(t) && sIdx < len(s) {
		if s[sIdx] == t[tIdx] {
			tIdx++
		}
		sIdx++
	}
	return len(t) - tIdx
}

package leetcode

func closeStrings(word1 string, word2 string) bool {
	ones := make(map[byte]int)
	twos := make(map[byte]int)

	for _, b := range word1 {
		ones[byte(b)]++
	}
	for _, b := range word2 {
		twos[byte(b)]++
	}

	if len(ones) != len(twos) {
		return false
	}

	for k := range ones {
		if _, ok := twos[k]; !ok {
			return false
		}
	}

	valueCount := make(map[int]int)
	for _, v := range ones {
		valueCount[v]++
	}

	for _, v := range twos {
		if n, ok := valueCount[v]; !ok || n <= 0 {
			return false
		}
		valueCount[v]--
	}

	return true
}

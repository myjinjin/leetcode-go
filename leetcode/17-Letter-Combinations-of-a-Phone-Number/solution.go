package leetcode

func letterCombinations(digits string) []string {
	combinations := []string{}
	phoneMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	if len(digits) == 0 {
		return []string{}
	}

	var backtrack func(input string, index int, curr string, combinations *[]string)
	backtrack = func(input string, index int, curr string, combinations *[]string) {
		if len(input) == index {
			*combinations = append(*combinations, curr)
			return
		}

		for _, char := range phoneMap[input[index]] {
			backtrack(input, index+1, curr+string(char), combinations)
		}
	}

	backtrack(digits, 0, "", &combinations)

	return combinations
}

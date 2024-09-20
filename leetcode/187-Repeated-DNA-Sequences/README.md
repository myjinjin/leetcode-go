# [187. Repeated DNA Sequences](https://leetcode.com/problems/repeated-dna-sequences/)

## Problem

The **DNA sequence** is composed of a series of nucleotides abbreviated as `'A'`, `'C'`, `'G'`, and `'T'`.

- For example, `"ACGAATTCCG"` is a **DNA sequence**.

When studying **DNA**, it is useful to identify repeated sequences within the DNA.

Given a string `s` that represents a *DNA sequence*, return all the `10`-letter-long sequences (substrings) that occur more than once in a DNA molecule. You may return the answer in **any order**.


Example 1:

```
Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Output: ["AAAAACCCCC","CCCCCAAAAA"]
```

Example 2:

```
Input: s = "AAAAAAAAAAAAA"
Output: ["AAAAAAAAAA"]
```

Constraints:

- `1 <= s.length <= 10^5`
- `s[i]` is either `'A'`, `'C'`, `'G'`, or `'T'`.

## Solution

```go
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	seen := make(map[string]int)
	result := make(map[string]bool)

	for i := 0; i <= len(s)-10; i++ {
		substr := s[i : i+10]
		seen[substr]++
		if seen[substr] > 1 {
			result[substr] = true
		}
	}

	var repeated []string
	for seq := range result {
		repeated = append(repeated, seq)
	}
	return repeated
}
```
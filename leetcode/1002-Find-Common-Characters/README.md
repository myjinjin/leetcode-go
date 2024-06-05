# [1002. Find Common Characters](https://leetcode.com/problems/find-common-characters/)

## Problem

Given a string array `words`, return an array of all characters that show up in all strings within the `words` (including duplicates). You may return the answer in any order.

 
Example 1:

```
Input: words = ["bella","label","roller"]
Output: ["e","l","l"]
```

Example 2:

```
Input: words = ["cool","lock","cook"]
Output: ["c","o"]
``` 

Constraints:

- `1 <= words.length <= 100`
- `1 <= words[i].length <= 100`
- `words[i]` consists of lowercase English letters.

## Solution

```go
func commonChars(words []string) []string {
    wordCount := make([]int, 26)
    for _, c := range words[0] {
        wordCount[int(c-'a')]++
    }

    for i:=1; i<len(words); i++ {
        word := words[i]
        tmpCount := make([]int, 26)
        for _, c := range word {
            tmpCount[int(c-'a')]++
        }
        for i:=0; i<len(wordCount); i++ {
            wordCount[i] = min(wordCount[i], tmpCount[i])
        }
    }

    var sb strings.Builder
    for i:=0; i<len(wordCount); i++ {
        count := wordCount[i]
        for count > 0 {
            sb.WriteByte(byte(i+int('a')))
            count--
        }
    }
    return strings.Split(sb.String(), "")
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```
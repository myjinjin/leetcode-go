# [273. Integer to English Words](https://leetcode.com/problems/integer-to-english-words/)

## Problem

Convert a non-negative integer `num` to its English words representation.


Example 1:

```
Input: num = 123
Output: "One Hundred Twenty Three"
```

Example 2:

```
Input: num = 12345
Output: "Twelve Thousand Three Hundred Forty Five"
```

Example 3:

```
Input: num = 1234567
Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
``` 

Constraints:

- `0 <= num <= 2^31 - 1`

## Solution

```go
func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	ones := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands := []string{"", "Thousand", "Million", "Billion"}

	var toWords func(num int) string
	toWords = func(num int) string {
		if num == 0 {
			return ""
		}
		if num < 20 {
			return ones[num] + " "
		}
		if num < 100 {
			return tens[num/10] + " " + toWords(num%10)
		}
		return ones[num/100] + " Hundred " + toWords(num%100)
	}

	var result []string
	for i := 0; num > 0; i++ {
		if num%1000 != 0 {
			words := strings.TrimSpace(toWords(num % 1000))
			if i > 0 {
				words += " " + thousands[i]
			}
			result = append([]string{words}, result...)
		}

		num /= 1000
	}

	return strings.Join(result, " ")
}
```
# [564. Find the Closest Palindrome](https://leetcode.com/problems/find-the-closest-palindrome/)

## Problem

Given a string `n` representing an integer, return the closest integer (not including itself), which is a palindrome. If there is a tie, return **the smaller one**.

The closest is defined as the absolute difference minimized between two integers.


Example 1:

```
Input: n = "123"
Output: "121"
```

Example 2:

```
Input: n = "1"
Output: "0"
Explanation: 0 and 2 are the closest palindromes but we return the smallest which is 0.
```

Constraints:

- `1 <= n.length <= 18`
- `n` consists of only digits.
- `n` does not have leading zeros.
- `n` is representing an integer in the range `[1, 10^18 - 1]`.

## Solution

```go
func nearestPalindromic(n string) string {
	if len(n) == 1 {
		return fmt.Sprintf("%d", (int(n[0]-'0') - 1))
	}

	num, _ := strconv.Atoi(n)
	leng := len(n)

	candidates := []int{}
	candidates = append(candidates, int(math.Pow10(leng-1)-1))
	candidates = append(candidates, int(math.Pow10(leng)+1))

	left, _ := strconv.ParseInt(n[:(leng+1)/2], 10, 64)
	for i := left - 1; i <= left+1; i++ {
		candidate := makePalindrome(int(i), leng%2 == 0)
		if candidate != num {
			candidates = append(candidates, candidate)
		}
	}

	closest := candidates[0]
	minDiff := abs(num - closest)
	for _, c := range candidates[1:] {
		diff := abs(num - c)
		if diff < minDiff || diff == minDiff && c < closest {
			minDiff = diff
			closest = c
		}
	}

	return fmt.Sprintf("%d", closest)
}

func makePalindrome(num int, isEven bool) int {
	palindrome := num
	if !isEven {
		num /= 10
	}
	for num > 0 {
		palindrome = palindrome*10 + num%10
		num /= 10
	}
	return palindrome
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
```
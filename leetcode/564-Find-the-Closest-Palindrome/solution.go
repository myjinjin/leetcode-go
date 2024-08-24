package leetcode

import (
	"fmt"
	"math"
	"strconv"
)

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

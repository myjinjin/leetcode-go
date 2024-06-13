package leetcode

import "math"

func guessNumber(n int) int {
	max := math.MaxInt32
	min := 1
	guessResult := guess(n)
	for min < max && guessResult != 0 {
		if guessResult == 1 {
			min = n
			n = min + (max-min)/2
			guessResult = guess(n)
		} else if guessResult == 0 {
			return n
		} else if guessResult == -1 {
			max = n
			n = min + (max-min)/2
			guessResult = guess(n)
		}
	}
	return n
}

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(n int) int {
	return 0
}

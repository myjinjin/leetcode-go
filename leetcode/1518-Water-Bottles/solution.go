package leetcode

func numWaterBottles(numBottles int, numExchange int) int {
	count := numBottles

	for numBottles >= numExchange {
		remainder := numBottles % numExchange
		numBottles = numBottles / numExchange
		count += numBottles
		numBottles += remainder
	}
	return count
}

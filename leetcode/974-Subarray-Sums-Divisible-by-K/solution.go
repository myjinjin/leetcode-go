package leetcode

func subarraysDivByK(nums []int, k int) int {
	remainderCount := make([]int, k)
	sum := 0
	count := 0
	remainderCount[0] = 1

	for _, num := range nums {
		sum = (sum + num%k + k) % k
		count += remainderCount[sum]
		remainderCount[sum]++
	}

	return count
}

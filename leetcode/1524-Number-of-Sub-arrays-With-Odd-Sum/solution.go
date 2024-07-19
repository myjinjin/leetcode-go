package leetcode

func numOfSubarrays(arr []int) int {
	currSum := 0
	sumArray := make([]int, 0, len(arr))
	for _, num := range arr {
		currSum = (currSum + num) % 2
		sumArray = append(sumArray, currSum)
	}

	oddSumCount, evenSumCount := 0, 0
	totalOddSumSubarrays := 0
	mod := 1000000007

	for _, sum := range sumArray {
		if sum%2 == 0 {
			evenSumCount++
			totalOddSumSubarrays += oddSumCount
		} else {
			oddSumCount++
			totalOddSumSubarrays += evenSumCount + 1
		}
		totalOddSumSubarrays %= mod
	}

	return totalOddSumSubarrays
}

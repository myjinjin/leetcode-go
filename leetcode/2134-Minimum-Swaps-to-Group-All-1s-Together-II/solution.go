package leetcode

func minSwaps(nums []int) int {
	n := len(nums)
	oneCount := 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			oneCount++
		}
	}

	expandedNums := append(nums, nums...)
	zeroCount := 0

	for i := 0; i < oneCount; i++ {
		if expandedNums[i] == 0 {
			zeroCount++
		}
	}

	minSwaps := zeroCount

	for i := oneCount; i < n+oneCount; i++ {
		if expandedNums[i-oneCount] == 0 {
			zeroCount--
		}
		if expandedNums[i] == 0 {
			zeroCount++
		}
		minSwaps = min(minSwaps, zeroCount)
	}

	return minSwaps
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 슬라이딩 윈도우 기법 사용
// 배열에 있는 모든 1의 개수를 세어 윈도우 크기를 결정한다
// 원형 배열을 다루기 위해 원래 배열을 두 번 이어붙인 새로운 배열을 만든다 -> 왜?
// 윈도우 크기만큼의 구간에서 0의 개수를 세어 필요한 스왑 횟수를 계산한다
// 윈도우를 한 칸씩 이동하면서 최소 스왑 횟수를 계산한다
// 윈도우를 한 칸씩 이동하면서 최소 스왑 횟수를 갱신한다

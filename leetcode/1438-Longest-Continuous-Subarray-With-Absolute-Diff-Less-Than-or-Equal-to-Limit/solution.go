package leetcode

func longestSubarray(nums []int, limit int) int {
	maxDequeue := make([]int, 0) // 부분 배열의 최댓값 관리
	minDequeue := make([]int, 0) // 부분 배열의 최솟값 관리

	count := 0
	left := 0 // left: 윈도우의 시작 인덱스, right: 윈도우의 끝 인덱스
	// right 포인터를 오른쪽으로 이동하면서 nums 배열의 원소를 하나씩 처리
	for right := 0; right < len(nums); right++ {
		// 현재 처리 중인 원소 nums[right]보다 작은 값들을 maxDequeue에서 제거
		// -> maxDequeue의 맨 앞에는 항상 현재 윈도우 내에서 가장 큰 값이 위치하게 된다
		for len(maxDequeue) > 0 && maxDequeue[len(maxDequeue)-1] < nums[right] {
			maxDequeue = maxDequeue[:len(maxDequeue)-1]
		}
		// 현재 처리 중인 원소 nums[right]를 maxDequeue의 끝에 추가
		maxDequeue = append(maxDequeue, nums[right])

		// 현재 처리 중인 원소 nums[right]보다 큰 값들을 minDequeue에서 제거
		// -> minDequeue의 맨 앞에는 항상 현재 윈도우 내에서 가장 작은 값이 위치하게 된다
		for len(minDequeue) > 0 && minDequeue[len(minDequeue)-1] > nums[right] {
			minDequeue = minDequeue[:len(minDequeue)-1]
		}
		// 현재 처리 중인 원소 nums[right]를 minDequeue의 끝에 추가
		minDequeue = append(minDequeue, nums[right])

		// maxDequeue의 맨 앞과 minDequeue의 맨 앞 값의 차이가 limit 보다 크다면, 현재 윈도우는 유효하지 않은 상태
		// 이때 left 포인터를 오른쪽으로 이동시키면서 윈도우를 축소한다
		for len(maxDequeue) > 0 && len(minDequeue) > 0 && maxDequeue[0]-minDequeue[0] > limit {
			// left 포인터가 maxDequeue의 맨 앞 값과 같다면, maxDequeue에서 해당 값을 제거한다
			if maxDequeue[0] == nums[left] {
				maxDequeue = maxDequeue[1:]
			}
			// left 포인터가 minDequeue의 맨 앞 값과 같다면, minDequeue에서 해당 값을 제거한다
			if minDequeue[0] == nums[left] {
				minDequeue = minDequeue[1:]
			}
			left++ // left 포인터를 오른쪽으로 한 칸 이동시킨다
		}
		// 현재 윈도우의 크기(right-left+1)와 이전에 찾은 최대 길이(count)를 비교하여 최대값 갱신
		count = max(count, right-left+1)
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

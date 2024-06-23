# [1438. Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit](https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/)

## Problem

Given an array of integers `nums` and an integer `limit`, return the size of the longest non-empty subarray such that the absolute difference between any two elements of this subarray is less than or equal to `limit`.


Example 1:

```
Input: nums = [8,2,4,7], limit = 4
Output: 2 
Explanation: All subarrays are: 
[8] with maximum absolute diff |8-8| = 0 <= 4.
[8,2] with maximum absolute diff |8-2| = 6 > 4. 
[8,2,4] with maximum absolute diff |8-2| = 6 > 4.
[8,2,4,7] with maximum absolute diff |8-2| = 6 > 4.
[2] with maximum absolute diff |2-2| = 0 <= 4.
[2,4] with maximum absolute diff |2-4| = 2 <= 4.
[2,4,7] with maximum absolute diff |2-7| = 5 > 4.
[4] with maximum absolute diff |4-4| = 0 <= 4.
[4,7] with maximum absolute diff |4-7| = 3 <= 4.
[7] with maximum absolute diff |7-7| = 0 <= 4. 
Therefore, the size of the longest subarray is 2.
```

Example 2:

```
Input: nums = [10,1,2,4,7,2], limit = 5
Output: 4 
Explanation: The subarray [2,4,7,2] is the longest since the maximum absolute diff is |2-7| = 5 <= 5.
```

Example 3:

```
Input: nums = [4,2,2,2,4,4,2,2], limit = 0
Output: 3
```

Constraints:

- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^9`
- `0 <= limit <= 10^9`

## Solution

```go
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
```
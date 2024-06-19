# [1482. Minimum Number of Days to Make m Bouquets](https://leetcode.com/problems/minimum-number-of-days-to-make-m-bouquets/)

## Problem

You are given an integer array `bloomDay`, an integer `m` and an integer `k`.

You want to make `m` bouquets. To make a bouquet, you need to use `k` adjacent flowers from the garden.

The garden consists of `n` flowers, the `ith` flower will bloom in the `bloomDay[i]` and then can be used in exactly one bouquet.

Return the minimum number of days you need to wait to be able to make `m` bouquets from the garden. If it is impossible to make m bouquets return `-1`.

 

Example 1:

```
Input: bloomDay = [1,10,3,10,2], m = 3, k = 1
Output: 3
Explanation: Let us see what happened in the first three days. x means flower bloomed and _ means flower did not bloom in the garden.
We need 3 bouquets each should contain 1 flower.
After day 1: [x, _, _, _, _]   // we can only make one bouquet.
After day 2: [x, _, _, _, x]   // we can only make two bouquets.
After day 3: [x, _, x, _, x]   // we can make 3 bouquets. The answer is 3.
```

Example 2:

```
Input: bloomDay = [1,10,3,10,2], m = 3, k = 2
Output: -1
Explanation: We need 3 bouquets each has 2 flowers, that means we need 6 flowers. We only have 5 flowers so it is impossible to get the needed bouquets and we return -1.
```

Example 3:

```
Input: bloomDay = [7,7,7,7,12,7,7], m = 2, k = 3
Output: 12
Explanation: We need 2 bouquets each should have 3 flowers.
Here is the garden after the 7 and 12 days:
After day 7: [x, x, x, x, _, x, x]
We can make one bouquet of the first three flowers that bloomed. We cannot make another bouquet from the last three flowers that bloomed because they are not adjacent.
After day 12: [x, x, x, x, x, x, x]
It is obvious that we can make two bouquets in different ways.
```

Constraints:

- `bloomDay.length == n`
- `1 <= n <= 10^5`
- `1 <= bloomDay[i] <= 10^9`
- `1 <= m <= 10^6`
- `1 <= k <= n`

## Solution

```go
func minDays(bloomDay []int, m int, k int) int {
	flowers := m * k
	if len(bloomDay) < flowers {
		return -1
	}
	result := -1

    // left: 가능한 최소 일수, right: 가능한 최대 일수
    // left가 right보다 커지면 반복문 종료
	left, right := 1, 1000000000
	for left <= right {
		mid := (left + right) / 2 // 중간값을 사용하여 m개의 꽃다발을 만들 수 있는지 확인
		length, bouquets := 0, 0
		for _, bloom := range bloomDay {
			if bloom <= mid {
				length++
				if length >= k {
					length = 0
					bouquets++
				}
			} else {
				length = 0
			}
		}
        // m개의 꽃다발을 만들 수 있는 경우 result 업데이트
        // right를 mid-1로 설정하여 탐색 범위 축소(최소 일수를 찾기 위해 왼쪽 범위 탐색)
		if bouquets >= m {
			result = mid
			right = mid - 1
		} else {
            // 꽃다발을 만들 수 없는 경우, left를 mid+1로 설정하여 탐색 범위 오른쪽으로 이동
			left = mid + 1
		}
	}
	return result
}
```
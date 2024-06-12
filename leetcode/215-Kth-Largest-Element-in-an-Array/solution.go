package leetcode

func findKthLargest(nums []int, k int) int {
	l, r := 0, len(nums)-1
	k = k - 1
	for l < r {
		i, j := quickSelect(l, r, nums)
		if k <= j {
			r = j
		} else if k >= i {
			l = i
		} else {
			break
		}
	}
	return nums[k]
}

func quickSelect(l, r int, nums []int) (int, int) {
	pivot := nums[(l+r)/2]
	for l <= r {
		for nums[r] < pivot {
			r--
		}
		for nums[l] > pivot {
			l++
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return l, r
}

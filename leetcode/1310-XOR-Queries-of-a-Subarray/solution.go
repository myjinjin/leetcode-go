package leetcode

func xorQueries(arr []int, queries [][]int) []int {
	answer := make([]int, len(queries))
	prefix := make([]int, len(arr)+1)
	for i := range arr {
		prefix[i+1] = prefix[i] ^ arr[i]
	}

	for i, q := range queries {
		left, right := q[0], q[1]
		answer[i] = prefix[right+1] ^ prefix[left]
	}
	return answer
}

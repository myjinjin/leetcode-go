package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = -1
		}
	}
	rowMin, rowMax := 0, m-1
	colMin, colMax := 0, n-1

	for head != nil {
		for j := colMin; j <= colMax; j++ {
			if head == nil {
				break
			}
			matrix[rowMin][j] = head.Val
			head = head.Next
		}
		rowMin++

		for i := rowMin; i <= rowMax; i++ {
			if head == nil {
				break
			}
			matrix[i][colMax] = head.Val
			head = head.Next
		}
		colMax--

		for j := colMax; j >= colMin; j-- {
			if head == nil {
				break
			}
			matrix[rowMax][j] = head.Val
			head = head.Next
		}
		rowMax--

		for i := rowMax; i >= rowMin; i-- {
			if head == nil {
				break
			}
			matrix[i][colMin] = head.Val
			head = head.Next
		}
		colMin++
	}

	return matrix
}

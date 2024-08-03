package leetcode

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	canReachIndex := make([]bool, n)
	canReachIndex[0] = true
	reachableCount := 0

	for i := 0; i < n; i++ {
		if i-minJump >= 0 && canReachIndex[i-minJump] {
			reachableCount++
		}
		if i-maxJump-1 >= 0 && canReachIndex[i-maxJump-1] {
			reachableCount--
		}

		if s[i] == '0' && reachableCount > 0 {
			canReachIndex[i] = true
		}
	}

	return canReachIndex[n-1]
}

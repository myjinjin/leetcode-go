package leetcode

import "math"

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	closestCost := math.MaxInt32

	var dfs func(index, currCost int)
	dfs = func(index, currCost int) {
		if index == len(toppingCosts) {
			if abs(currCost-target) < abs(closestCost-target) ||
				(abs(currCost-target) == abs(closestCost-target) && currCost < closestCost) {
				closestCost = currCost
			}
			return
		}

		dfs(index+1, currCost)
		dfs(index+1, currCost+toppingCosts[index])
		dfs(index+1, currCost+2*toppingCosts[index])
	}

	for _, baseCost := range baseCosts {
		dfs(0, baseCost)
	}

	return closestCost
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

package leetcode

import (
	"sort"
)

func maximumTotalDamage(power []int) int64 {
	damageFrequency := make(map[int]int)
	for _, damage := range power {
		damageFrequency[damage]++
	}

	damages := make([]int, 0, len(damageFrequency))
	for damage := range damageFrequency {
		damages = append(damages, damage)
	}
	sort.Ints(damages)

	n := len(damages)
	maxDamage := make([]int64, n)
	maxDamage[0] = int64(damages[0] * damageFrequency[damages[0]])

	for i := 1; i < n; i++ {
		currDamage := damages[i]
		currDamageSum := int64(currDamage * damageFrequency[currDamage])

		maxDamage[i] = maxDamage[i-1]
		prev := i - 1
		for prev >= 0 && (damages[prev] == currDamage-1 ||
			damages[prev] == currDamage-2 ||
			damages[prev] == currDamage+1 ||
			damages[prev] == currDamage+2) {
			prev--
		}

		if prev >= 0 {
			maxDamage[i] = max(maxDamage[i], maxDamage[prev]+currDamageSum)
		} else {
			maxDamage[i] = max(maxDamage[i], currDamageSum)
		}
	}

	return maxDamage[n-1]
}

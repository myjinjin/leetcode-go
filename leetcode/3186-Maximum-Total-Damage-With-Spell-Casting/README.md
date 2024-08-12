# [3186. Maximum Total Damage With Spell Casting](https://leetcode.com/problems/maximum-total-damage-with-spell-casting/)

## Problem

A magician has various spells.

You are given an array `power`, where each element represents the damage of a spell. Multiple spells can have the same damage value.

It is a known fact that if a magician decides to cast a spell with a damage of `power[i]`, they **cannot** cast any spell with a damage of `power[i] - 2`, `power[i] - 1`, `power[i] + 1`, or `power[i] + 2`.

Each spell can be cast **only once**.

Return the maximum possible total damage that a magician can cast.


Example 1:

```
Input: power = [1,1,3,4]

Output: 6

Explanation:

The maximum possible damage of 6 is produced by casting spells 0, 1, 3 with damage 1, 1, 4.
```

Example 2:

```
Input: power = [7,1,6,6]

Output: 13

Explanation:

The maximum possible damage of 13 is produced by casting spells 1, 2, 3 with damage 1, 6, 6.
```
 

Constraints:

- `1 <= power.length <= 10^5`
- `1 <= power[i] <= 10^9`


## Solution

```go
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
```
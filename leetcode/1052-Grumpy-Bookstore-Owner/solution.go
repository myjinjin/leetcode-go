package leetcode

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n := len(customers)
	satisfied := 0

	if n == minutes {
		for i := 0; i < n; i++ {
			satisfied += customers[i]
		}
		return satisfied
	}

	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			satisfied += customers[i]
			customers[i] = 0
		}
	}
	maxMinutesSatisfiedSum := 0
	minutesSatisfiedSum := 0
	for i := 0; i < minutes; i++ {
		minutesSatisfiedSum += customers[i]
	}
	maxMinutesSatisfiedSum = minutesSatisfiedSum

	for i := 0; i < n-minutes; i++ {
		sub := customers[i]
		add := customers[i+minutes]
		minutesSatisfiedSum = minutesSatisfiedSum - sub + add
		maxMinutesSatisfiedSum = max(maxMinutesSatisfiedSum, minutesSatisfiedSum)
	}

	return satisfied + maxMinutesSatisfiedSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

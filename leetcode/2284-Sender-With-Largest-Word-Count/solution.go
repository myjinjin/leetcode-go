package leetcode

import "strings"

func largestWordCount(messages []string, senders []string) string {
	wordCount := make(map[string]int)
	largestCount := 0
	largestSender := ""
	for i := 0; i < len(messages); i++ {
		wordCount[senders[i]] += len(strings.Split(messages[i], " "))
		if largestCount < wordCount[senders[i]] {
			largestCount = wordCount[senders[i]]
			largestSender = senders[i]
		} else if largestCount == wordCount[senders[i]] {
			if largestSender < senders[i] {
				largestSender = senders[i]
			}
		}
	}
	return largestSender
}

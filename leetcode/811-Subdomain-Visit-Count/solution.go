package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	visits := make(map[string]int)

	for _, d := range cpdomains {
		splited := strings.Split(d, " ")
		numStr, domain := splited[0], splited[1]
		num, _ := strconv.Atoi(numStr)
		domains := strings.Split(domain, ".")
		for i := 0; i < len(domains); i++ {
			visits[strings.Join(domains[i:], ".")] += num
		}
	}

	result := make([]string, 0, len(visits))
	for k, v := range visits {
		result = append(result, fmt.Sprintf("%d %s", v, k))
	}
	return result
}

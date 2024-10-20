package leetcode

import "strings"

func defangIPaddr(address string) string {
	address = strings.ReplaceAll(address, ".", "[.]")
	return address
}

package leetcode

import "strings"

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	ones := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	thousands := []string{"", "Thousand", "Million", "Billion"}

	var toWords func(num int) string
	toWords = func(num int) string {
		if num == 0 {
			return ""
		}
		if num < 20 {
			return ones[num] + " "
		}
		if num < 100 {
			return tens[num/10] + " " + toWords(num%10)
		}
		return ones[num/100] + " Hundred " + toWords(num%100)
	}

	var result []string
	for i := 0; num > 0; i++ {
		if num%1000 != 0 {
			words := strings.TrimSpace(toWords(num % 1000))
			if i > 0 {
				words += " " + thousands[i]
			}
			result = append([]string{words}, result...)
		}

		num /= 1000
	}

	return strings.Join(result, " ")
}

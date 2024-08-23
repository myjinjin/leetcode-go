# [592. Fraction Addition and Subtraction](https://leetcode.com/problems/fraction-addition-and-subtraction/)

## Problem

Given a string `expression` representing an expression of fraction addition and subtraction, return the calculation result in string format.

The final result should be an irreducible fraction. If your final result is an integer, change it to the format of a fraction that has a denominator `1`. So in this case, `2` should be converted to `2/1`.


Example 1:

```
Input: expression = "-1/2+1/2"
Output: "0/1"
```

Example 2:

```
Input: expression = "-1/2+1/2+1/3"
Output: "1/3"
```

Example 3:

```
Input: expression = "1/3-1/2"
Output: "-1/6"
``` 


Constraints:

- The input string only contains `'0'` to `'9'`, `'/'`, `'+'` and `'-'`. So does the output.
- Each fraction (input and output) has the format `±numerator/denominator`. If the first input fraction or the output is positive, then `'+'` will be omitted.
- The input only contains valid irreducible fractions, where the numerator and denominator of each fraction will always be in the range `[1, 10]`. If the denominator is `1`, it means this fraction is actually an integer in a fraction format defined above.
- The number of given fractions will be in the range `[1, 10]`.
- The numerator and denominator of the final result are guaranteed to be valid and in the range of **32-bit** int.

## Solution

```go
func fractionAddition(expression string) string {
	i := 0
	numerator, denominator := 0, 1

	for i < len(expression) {
		sign := 1
		if expression[i] == '-' || expression[i] == '+' {
			if expression[i] == '-' {
				sign = -1
			}
			i++
		}

		var num int
		for i < len(expression) && expression[i] >= '0' && expression[i] <= '9' {
			num *= 10
			num += int(expression[i] - '0')
			i++
		}
		num *= sign
		i++

		var den int
		for i < len(expression) && expression[i] >= '0' && expression[i] <= '9' {
			den *= 10
			den += int(expression[i] - '0')
			i++
		}

		numerator = num*denominator + numerator*den
		denominator *= den
	}

	g := gcd(numerator, denominator)
	numerator /= g
	denominator /= g
	if denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	}

	return fmt.Sprintf("%d/%d", numerator, denominator)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
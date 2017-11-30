// UVa 290 - Palindroms <---> smordnilaP

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumber(digit rune) int {
	if digit >= '0' && digit <= '9' {
		return int(digit - '0')
	}
	return int(digit-'A') + 10
}

func minBase(num string) int {
	base := 1
	for _, d := range num {
		if tmp := getNumber(d); tmp > base {
			base = tmp
		}
	}
	return base + 1
}

func toDigits(num string) []int {
	digits := make([]int, len(num))
	for i, d := range num {
		digits[i] = getNumber(d)
	}
	return digits
}

func reverse(digits []int) []int {
	l := len(digits)
	ret := make([]int, l)
	for i := range ret {
		ret[i] = digits[l-1-i]
	}
	return ret
}

func toBase(n, base int) []int {
	var digits []int
	for ; n > 0; n /= base {
		digits = append([]int{n % base}, digits...)
	}
	return digits
}

func isPalindrome(digits []int) bool {
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}

func toNumber(digits []int, base int) int {
	var total int
	for _, d := range digits {
		total = total*base + d
	}
	return total
}

func do(num, base int) int {
	for step := 1; ; step++ {
		sum := num + toNumber(reverse(toBase(num, base)), base)
		if isPalindrome(toBase(sum, base)) {
			return step
		}
		num = sum
	}
}

func solve(num string) string {
	var steps []string
	min := minBase(num)
	for base := 15; base >= 2; base-- {
		if base >= min {
			steps = append(steps, strconv.Itoa(do(toNumber(toDigits(num), base), base)))
		} else {
			steps = append(steps, "?")
		}
	}
	return strings.Join(steps, " ")
}

func main() {
	in, _ := os.Open("290.in")
	defer in.Close()
	out, _ := os.Create("290.out")
	defer out.Close()

	var num string
	for {
		if _, err := fmt.Fscanf(in, "%s", &num); err != nil {
			break
		}
		fmt.Fprintln(out, solve(num))
	}
}

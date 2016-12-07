// UVa 389 - Basically Speaking

package main

import (
	"fmt"
	"math"
	"os"
)

func toNum(digit rune) int {
	if digit >= '0' && digit <= '9' {
		return int(digit - '0')
	}
	if digit >= 'A' && digit <= 'Z' {
		return int(digit-'A') + 10
	}
	return -1
}

func base10(num string, base int) int {
	len := len(num)
	total := 0
	for i, d := range num {
		digit := toNum(d)
		if digit >= base {
			return -1
		}
		total += digit * int(math.Pow(float64(base), float64(len-i-1)))
	}
	return total
}

func toStr(d int) string {
	if d <= 9 {
		return string('0' + d)
	} else {
		return string('A' + d - 10)
	}
}

func baseN(num int, base int) string {
	var number string
	for num > 0 {
		digit := num % base
		number = toStr(digit) + number
		num = (num - digit) / base
	}
	return number
}

func main() {
	in, _ := os.Open("389.in")
	defer in.Close()
	out, _ := os.Create("389.out")
	defer out.Close()

	var number string
	var b1, b2 int
	for {
		if _, err := fmt.Fscanf(in, "%s%d%d", &number, &b1, &b2); err != nil {
			break
		}
		bt := base10(number, b1)
		if num := baseN(bt, b2); len(num) > 7 {
			fmt.Fprintln(out, "  ERROR")
		} else {
			fmt.Fprintf(out, "%7s\n", num)
		}
	}
}

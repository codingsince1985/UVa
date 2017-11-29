// UVa 343 - What Base Is This?

package main

import (
	"fmt"
	"os"
)

var numbers [][2]string

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

func base10(num string, base int) int {
	var total int
	for _, d := range num {
		total = total*base + getNumber(d)
	}
	return total
}

func main() {
	in, _ := os.Open("343.in")
	defer in.Close()
	out, _ := os.Create("343.out")
	defer out.Close()

	var a, b string
	for {
		if _, err := fmt.Fscanf(in, "%s%s", &a, &b); err != nil {
			break
		}
		numbers = append(numbers, [2]string{a, b})
	}

here:
	for _, nums := range numbers {
		a1, a2 := nums[0], nums[1]
		b1, b2 := minBase(a1), minBase(a2)
		for i := b1; i <= 36; i++ {
			for j, num1 := b2, base10(a1, i); j <= 36; j++ {
				if num1 == base10(a2, j) {
					fmt.Fprintf(out, "%s (base %d) = %s (base %d)\n", a1, i, a2, j)
					continue here
				}
			}
		}
		fmt.Fprintf(out, "%s is not equal to %s in any base 2..36\n", a1, a2)
	}
}

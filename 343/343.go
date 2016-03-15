// UVa 343 - What Base Is This?

package main

import (
	"fmt"
	"math"
	"os"
)

var numbers [][2]string

func getNumber(digit rune) int {
	if digit >= '0' && digit <= '9' {
		return int(digit - '0')
	}
	if digit >= 'A' && digit <= 'Z' {
		return int(digit-'A') + 10
	}
	return -1
}

func minBase(num string) int {
	var base = 1
	for _, d := range num {
		tmp := getNumber(d)
		if tmp > base {
			base = tmp
		}
	}
	return base + 1
}

func base10(number string, base int) int {
	total := 0
	len := len(number)
	for i, d := range number {
		total += getNumber(d) * int(math.Pow(float64(base), float64(len-i-1)))
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

	for _, nums := range numbers {
		a1, a2 := nums[0], nums[1]
		b1, b2 := minBase(a1), minBase(a2)
		found := false
		for i := b1; i <= 36 && !found; i++ {
			num1 := base10(a1, i)
			for j := b2; j <= 36; j++ {
				num2 := base10(a2, j)
				if num1 == num2 {
					fmt.Fprintf(out, "%s (base %d) = %s (base %d)\n", a1, i, a2, j)
					found = true
					break
				}
			}
		}
		if !found {
			fmt.Fprintf(out, "%s is not equal to %s in any base 2..36\n", a1, a2)
		}
	}
}

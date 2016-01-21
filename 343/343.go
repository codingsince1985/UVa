// UVa 343 - What Base Is This?

package main

import (
	"fmt"
	"math"
	"os"
)

var numbers [][2]string

func prepare() {
	in, _ := os.Open("343.in")
	defer in.Close()

	numbers = make([][2]string, 0)
	var i = -1
	var a, b string
	i, _ = fmt.Fscanf(in, "%s%s", &a, &b)
	for ; i != 0; {
		numbers = append(numbers, [...]string{a, b})
		i, _ = fmt.Fscanf(in, "%s%s", &a, &b)
	}
}

func getNumber(d rune) int {
	digit := int(d)
	if digit >= 48 && digit <= 57 {
		return digit - 48
	} else if digit >= 65 && digit <= 90 {
		return digit - 55
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
		total += getNumber(d) * int(math.Pow(float64(base), float64(len - i - 1)))
	}
	return total
}

func do() {
	out, _ := os.Create("343.out")
	defer out.Close()

	for _, nums := range numbers {
		a1, a2 := nums[0], nums[1]
		b1, b2 := minBase(a1), minBase(a2)
		found := false
		for i := b1; i <= 36 && !found; i ++ {
			num1 := base10(a1, i)
			for j := b2; j <= 36; j ++ {
				num2 := base10(a2, j)
				if num1 == num2 {
					fmt.Fprintf(out, "%s (base %d) = %s (base %d)\n", a1, i, a2, j)
					found = true;
					break;
				}
			}
		}
		if !found {
			fmt.Fprintf(out, "%s is not equal to %s in any base 2..36\n", a1, a2)
		}
	}
}

func main() {
	prepare()
	do()
}
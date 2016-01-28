// UVa 355 - The Bases Are Loaded

package main

import (
	"fmt"
	"math"
	"os"
)

type Number struct {
	b1, b2 int
	num    string
}

var nums []Number

func prepare() {
	in, _ := os.Open("355.in")
	defer in.Close()
	nums = nil

	var b1, b2 int
	var num string
	i, _ := fmt.Fscanf(in, "%d%d%s", &b1, &b2, &num)
	for i > 0 {
		nums = append(nums, Number{b1, b2, num})
		i, _ = fmt.Fscanf(in, "%d%d%s", &b1, &b2, &num)
	}
}

func toStr(d int) string {
	if d <= 9 {
		return string('0' + d)
	}
	return string('A' - 10 + d)
}

func toNum(digit rune) int {
	if digit >= '0' && digit <= '9' {
		return int(digit - '0')
	} else if digit >= 'A' && digit <= 'Z' {
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

func baseN(num int, base int) string {
	var number string
	for num > 0 {
		digit := num % base
		number = toStr(digit) + number
		num = (num - digit) / base
	}
	return number
}

func do() {
	out, _ := os.Create("355.out")
	defer out.Close()

	for _, number := range nums {
		bt := base10(number.num, number.b1)
		if bt == -1 {
			fmt.Fprintf(out, "%s is an illegal base %d number\n", number.num, number.b1)
		} else {
			fmt.Fprintf(out, "%s base %d = %s base %d\n", number.num, number.b1, baseN(bt, number.b2), number.b2)
		}
	}
}

func main() {
	prepare()
	do()
}

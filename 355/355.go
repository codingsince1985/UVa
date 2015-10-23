// UVa 355 - The Bases Are Loaded

package main

import (
	"fmt"
	"os"
	"math"
)

type Number struct {
	b1, b2 int
	num    string
}

var nums []Number
var out *os.File

func prepare() {
	in, _ := os.Open("355.in")
	nums = make([]Number, 0)

	var b1, b2 int
	var num string
	i, _ := fmt.Fscanf(in, "%d %d %s", &b1, &b2, &num)
	for ; i > 0; {
		nums = append(nums, Number{b1, b2, num})
		i, _ = fmt.Fscanf(in, "%d %d %s", &b1, &b2, &num)
	}

	out, _ = os.Create("355.out")
}

func toStr(d int) string {
	if d <= 9 {
		return string(48 + d)
	} else {
		return string(55 + d)
	}
}

func toNum(digit rune) int {
	d := int(digit)
	if d >= 48 && d <= 57 {
		return d - 48
	} else if d >= 65 && d <= 90 {
		return d - 55
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
		total += digit * int(math.Pow(float64(base), float64(len - i - 1)))
	}
	return total
}

func baseN(num int, base int) string {
	var number string
	for num > 0 {
		digit := num % base
		number += toStr(digit)
		num = (num - digit) / base
	}
	return number
}

func do() {
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
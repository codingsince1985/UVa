// UVa 355 - The Bases Are Loaded

package main

import (
	"fmt"
	"os"
)

type number struct {
	b1, b2 int
	num    string
}

var nums []number

func toStr(d int) string {
	if d <= 9 {
		return string('0' + d)
	}
	return string('A' - 10 + d)
}

func getNumber(digit rune) int {
	if digit >= '0' && digit <= '9' {
		return int(digit - '0')
	}
	if digit >= 'A' && digit <= 'Z' {
		return int(digit-'A') + 10
	}
	return -1
}

func base10(num string, base int) int {
	var total int
	for _, d := range num {
		digit := getNumber(d)
		if digit >= base {
			return -1
		}
		total = total*base + digit
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

func main() {
	in, _ := os.Open("355.in")
	defer in.Close()
	out, _ := os.Create("355.out")
	defer out.Close()

	var num number
	for {
		if _, err := fmt.Fscanf(in, "%d%d%s", &num.b1, &num.b2, &num.num); err != nil {
			break
		}
		nums = append(nums, num)
	}

	for _, number := range nums {
		if bt := base10(number.num, number.b1); bt == -1 {
			fmt.Fprintf(out, "%s is an illegal base %d number\n", number.num, number.b1)
		} else {
			fmt.Fprintf(out, "%s base %d = %s base %d\n", number.num, number.b1, baseN(bt, number.b2), number.b2)
		}
	}
}

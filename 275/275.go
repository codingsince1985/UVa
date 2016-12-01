// UVa 275 - Expanding Fractions

package main

import (
	"fmt"
	"os"
	"strconv"
)

func divide(n, d int) (int, string) {
	cache := make(map[int]int)
	digits := []byte{'.'}
	var pre int
	var ok bool
	for n != 0 {
		if pre, ok = cache[n]; ok {
			break
		}
		cache[n] = len(digits)
		n *= 10
		digits = append(digits, strconv.Itoa(n / d)[0])
		n %= d
	}
	if n == 0 {
		return 0, string(digits)
	} else {
		return len(digits) - pre, string(digits)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func output(out *os.File, length int, digits string) {
	for i := 0; i <= len(digits)/50; i++ {
		segment := digits[i*50 : min((i+1)*50, len(digits))]
		if len(segment) > 0 {
			fmt.Fprintln(out, segment)
		}
	}
	if length == 0 {
		fmt.Fprintln(out, "This expansion terminates.\n")
	} else {
		fmt.Fprintf(out, "The last %d digits repeat forever.\n\n", length)
	}
}

func main() {
	in, _ := os.Open("275.in")
	defer in.Close()
	out, _ := os.Create("275.out")
	defer out.Close()

	var n, d int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &d); n == 0 && d == 0 {
			break
		}
		length, digits := divide(n, d)
		output(out, length, digits)
	}
}

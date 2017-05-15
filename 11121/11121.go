// UVa 11121 - Base -2

package main

import (
	"fmt"
	"os"
)

const base = -2

func solve(n int) string {
	if n == 0 {
		return "0"
	}
	var digits []byte
	for n != 0 {
		digit := n % base
		n /= base
		if digit == -1 {
			digit = 1
			n++
		}
		digits = append([]byte{byte(digit + '0')}, digits...)
	}
	return string(digits)
}

func main() {
	in, _ := os.Open("11121.in")
	defer in.Close()
	out, _ := os.Create("11121.out")
	defer out.Close()

	var kase, n int
	fmt.Fscanf(in, "%d", &kase)
	for i := 1; i <= kase; i++ {
		fmt.Fscanf(in, "%d", &n)
		fmt.Fprintf(out, "Case #%d: %s\n", i, solve(n))
	}
}

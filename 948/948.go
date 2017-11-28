// UVa 948 - Fibonaccimal Base

package main

import (
	"fmt"
	"os"
	"strings"
)

const max = 39

var fibonacci = func() []int {
	f := make([]int, max+1)
	f[0], f[1] = 0, 1
	for i := 2; i <= max; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}()

func solve(dec int) string {
	var fib string
	for i := max; i > 1; i-- {
		if dec >= fibonacci[i] {
			dec -= fibonacci[i]
			fib += "1"
		} else {
			fib += "0"
		}
	}
	return strings.TrimLeft(fib, "0")
}

func main() {
	in, _ := os.Open("948.in")
	defer in.Close()
	out, _ := os.Create("948.out")
	defer out.Close()

	var n, dec int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d", &dec)
		fmt.Fprintf(out, "%d = %s (fib)\n", dec, solve(dec))
	}
}

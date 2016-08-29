// UVa 636 - Squares (III)

package main

import (
	"fmt"
	"math"
	"os"
)

func isValid(n, i int) bool {
	for ; n > 0; n /= 10 {
		if n%10 >= i {
			return false
		}
	}
	return true
}

func calc(n, base int) int {
	var num int
	b := 1
	for ; n > 0; n /= 10 {
		num += (n % 10) * b
		b *= base
	}
	return num
}

func isSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n) + .5))
	return sqrt*sqrt == n
}

func solve(n int) (i int) {
	for i = 2; i < 100; i++ {
		if isValid(n, i) {
			break
		}
	}
	for ; i < 100; i++ {
		if isSquare(calc(n, i)) {
			return
		}
	}
	return
}

func main() {
	in, _ := os.Open("636.in")
	defer in.Close()
	out, _ := os.Create("636.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintln(out, solve(n))
	}
}

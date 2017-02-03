// UVa 993 - Product of digits

package main

import (
	"fmt"
	"os"
)

func solve(n int) int {
	if n == 1 {
		return 1
	}
	var factors int
	for i, base := 9, 1; i > 1 && n > 1; i-- {
		for ; n%i == 0; n, base = n/i, base*10 {
			factors = i*base + factors
		}
	}
	if n != 1 {
		return -1
	}
	return factors
}

func main() {
	in, _ := os.Open("993.in")
	defer in.Close()
	out, _ := os.Create("993.out")
	defer out.Close()

	var kase, n int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d", &n)
		fmt.Fprintln(out, solve(n))
	}
}

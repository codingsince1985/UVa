// UVa 10105 - Polynomial Coefficients

package main

import (
	"fmt"
	"os"
)

func factorial(n int) int {
	fact := 1
	for n > 0 {
		fact *= n
		n--
	}
	return fact
}

func main() {
	in, _ := os.Open("10105.in")
	defer in.Close()
	out, _ := os.Create("10105.out")
	defer out.Close()

	var n, k, nk int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &k); err != nil {
			break
		}
		// n! / (n1! * n2! * ... * nk!)
		coefficient := factorial(n)
		for k > 0 {
			k--
			fmt.Fscanf(in, "%d", &nk)
			coefficient /= factorial(nk)
		}
		fmt.Fprintln(out, coefficient)
	}
}

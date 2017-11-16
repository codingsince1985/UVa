// UVa 10302 - Summation of Polynomials

package main

import (
	"fmt"
	"math/big"
	"os"
)

// Σx^3 = (n^2) * (n+1)^2 / 4;

func main() {
	in, _ := os.Open("10302.in")
	defer in.Close()
	out, _ := os.Create("10302.out")
	defer out.Close()

	var n int64
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		p1 := big.NewInt(n)
		p2 := big.NewInt(n + 1)
		p1.Mul(p1, p1)
		p2.Mul(p2, p2)
		fmt.Fprintln(out, p1.Mul(p1, p2).Div(p1, big.NewInt(4)))
	}
}

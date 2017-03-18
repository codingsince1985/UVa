// UVa 623 - 500!

package main

import (
	"fmt"
	"math/big"
	"os"
)

func factorial(n int64) *big.Int {
	fact := big.NewInt(1)
	for ; n > 1; n-- {
		fact.Mul(fact, big.NewInt(n))
	}
	return fact
}

func main() {
	in, _ := os.Open("623.in")
	defer in.Close()
	out, _ := os.Create("623.out")
	defer out.Close()

	var n int64
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintf(out, "%d!\n", n)
		fmt.Fprintln(out, factorial(n))
	}
}

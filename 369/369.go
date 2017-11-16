// UVa 369 - Combinations

package main

import (
	"fmt"
	"math/big"
	"os"
)

func f(n int) *big.Int {
	r := big.NewInt(1)
	for i := 2; i <= n; i++ {
		r.Mul(r, big.NewInt(int64(i)))
	}
	return r
}

func c(n, m int) *big.Int {
	r := f(n)
	s := f(n - m)
	t := f(m)
	return r.Div(r, s).Div(r, t)
}

func main() {
	in, _ := os.Open("369.in")
	defer in.Close()
	out, _ := os.Create("369.out")
	defer out.Close()

	var n, m int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		fmt.Fprintf(out, "%d things taken %d at a time is %v exactly.\n", n, m, c(n, m))
	}
}

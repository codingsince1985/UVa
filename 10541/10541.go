// UVa 10541 - Stripe

package main

import (
	"fmt"
	"math/big"
	"os"
)

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

func factorial(a int) *big.Int {
	f := big.NewInt(1)
	for ; a > 1; a-- {
		f.Mul(f, big.NewInt(int64(a)))
	}
	return f
}

func solve(n, k int) *big.Int {
	switch {
	case n == 0:
		if k == 0 {
			return one
		}
		return zero
	case k > n:
		return zero
	case k == n:
		return one
	default:
		fn, fk := factorial(n), factorial(k)
		return fn.Div(fn, fk.Mul(fk, factorial(n-k)))
	}
}

func main() {
	in, _ := os.Open("10541.in")
	defer in.Close()
	out, _ := os.Create("10541.out")
	defer out.Close()

	var t, n, k, black int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d%d", &n, &k)
		var lenBlack int
		for i := 0; i < k; i++ {
			fmt.Fscanf(in, "%d", &black)
			lenBlack += black
		}
		fmt.Fprintln(out, solve(n-lenBlack+1, k))
	}
}

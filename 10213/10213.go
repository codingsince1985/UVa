// UVa 10213 - How Many Pieces of Land ?

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

func factorial(n int64) *big.Int {
	fact := big.NewInt(1)
	for ; n > 1; n-- {
		fact.Mul(fact, big.NewInt(n))
	}
	return fact
}

func combination(n, r int64) *big.Int {
	if n < r {
		return zero
	}
	if n == r {
		return one
	}
	n1 := factorial(n)
	r1 := factorial(r)
	nr := factorial(n - r)
	tmp := n1
	tmp.Div(tmp, r1)
	tmp.Div(tmp, nr)
	return tmp
}

func solve(n int64) *big.Int {
	c1 := combination(n, 4)
	c2 := combination(n, 2)
	tmp := big.NewInt(1)
	tmp.Add(tmp, c1)
	tmp.Add(tmp, c2)
	return tmp
}

func main() {
	in, _ := os.Open("10213.in")
	defer in.Close()
	out, _ := os.Create("10213.out")
	defer out.Close()

	var kase, n int64
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d", &n)
		fmt.Fprintln(out, solve(n))
	}
}

// UVa 10007 - Count the Trees

package main

import (
	"fmt"
	"math/big"
	"os"
)

const MAX = 301

var catalan, fact [MAX]big.Int

func prepare() {
	catalan[0].SetInt64(1)
	for i := 1; i < MAX; i++ {
		catalan[i].Mul(&catalan[i-1], big.NewInt(int64(2*(2*i-1))))
		catalan[i].Div(&catalan[i], big.NewInt(int64(i+1)))
	}

	fact[1].SetInt64(1)
	for i := 2; i < MAX; i++ {
		fact[i].Mul(&fact[i-1], big.NewInt(int64(i)))
	}
}

func main() {
	prepare()

	in, _ := os.Open("10007.in")
	defer in.Close()
	out, _ := os.Create("10007.out")
	defer out.Close()

	var tmp big.Int
	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		tmp.Mul(&fact[n], &catalan[n])
		fmt.Fprintf(out, "%v\n", &tmp)
	}
}

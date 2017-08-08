// UVa 10007 - Count the Trees

package main

import (
	"fmt"
	"math/big"
	"os"
)

const max = 301

var catalan, fact = func() ([max]big.Int, [max]big.Int) {
	var catalan, fact [max]big.Int
	catalan[0].SetInt64(1)
	for i := int64(1); i < max; i++ {
		catalan[i].Mul(&catalan[i-1], big.NewInt(2*(2*i-1)))
		catalan[i].Div(&catalan[i], big.NewInt(i+1))
	}

	fact[1].SetInt64(1)
	for i := int64(2); i < max; i++ {
		fact[i].Mul(&fact[i-1], big.NewInt(i))
	}
	return catalan, fact
}()

func main() {
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

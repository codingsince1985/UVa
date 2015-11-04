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
		catalan[i].Mul(&catalan[i - 1], big.NewInt(int64(2 * (2 * i - 1))))
		catalan[i].Div(&catalan[i], big.NewInt(int64(i + 1)))
	}

	fact[1].SetInt64(1)
	for i := 2; i < MAX; i++ {
		fact[i].Mul(&fact[i - 1], big.NewInt(int64(i)))
	}
}

func do() {
	in, _ := os.Open("10007.in")
	out, _ := os.Create("10007.out")
	var tmp big.Int
	var n int
	for {
		fmt.Fscanf(in, "%d", &n)
		if n == 0 {
			break;
		}
		tmp.Mul(&fact[n], &catalan[n])
		fmt.Fprintf(out, "%v\n", &tmp)
	}
	in.Close()
	out.Close()
}
func main() {
	prepare()
	do()
}
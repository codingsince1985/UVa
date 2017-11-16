// UVa 11879 - Multiple of 17

package main

import (
	"fmt"
	"math/big"
	"os"
)

var (
	five      = big.NewInt(5)
	ten       = big.NewInt(10)
	seventeen = big.NewInt(17)
)

func solve(n string) bool {
	var num, last big.Int
	num.SetString(n, 10)
	num.DivMod(&num, ten, &last)
	num.Sub(&num, last.Mul(&last, five)).DivMod(&num, seventeen, &last)
	return last.Int64() == 0
}

func main() {
	in, _ := os.Open("11879.in")
	defer in.Close()
	out, _ := os.Create("11879.out")
	defer out.Close()

	var n string
	for {
		if fmt.Fscanf(in, "%s", &n); n == "0" {
			break
		}
		if solve(n) {
			fmt.Fprintln(out, 1)
		} else {
			fmt.Fprintln(out, 0)
		}
	}
}

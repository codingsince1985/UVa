// UVa 495 - Fibonacci Freeze

package main

import (
	"fmt"
	"math/big"
	"os"
)

var f [5001]*big.Int

func prepare() {
	f[0] = big.NewInt(0)
	f[1] = big.NewInt(1)
	var tmp *big.Int
	for i := 2; i <= 5000; i ++ {
		tmp = big.NewInt(0)
		tmp.Add(f[i - 1], f[i - 2])
		f[i] = tmp
	}
}

func main() {
	prepare()

	in, _ := os.Open("495.in")
	out, _ := os.Create("495.out")
	var n int
	for {
		_, err := fmt.Fscanf(in, "%d", &n)
		if err != nil {
			break;
		}
		fmt.Fprintf(out, "The Fibonacci number for %d is ", n)
		fmt.Fprintf(out, "%v\n", f[n])
	}
}

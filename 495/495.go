// UVa 495 - Fibonacci Freeze

package main

import (
	"fmt"
	"math/big"
	"os"
)

const MAX = 5001

var f [MAX]big.Int

func prepare() {
	f[0].SetInt64(0)
	f[1].SetInt64(1)
	for i := 2; i < MAX; i ++ {
		f[i].Add(&f[i - 1], &f[i - 2])
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
		fmt.Fprintf(out, "The Fibonacci number for %d is %v\n", n, &f[n])
	}
	in.Close()
	out.Close()
}

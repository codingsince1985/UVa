// UVa 495 - Fibonacci Freeze

package main

import (
	"fmt"
	"math/big"
	"os"
)

const max = 5001

var f = func() [max]big.Int {
	var f [max]big.Int
	f[0].SetInt64(0)
	f[1].SetInt64(1)
	for i := 2; i < max; i++ {
		f[i].Add(&f[i-1], &f[i-2])
	}
	return f
}()

func main() {
	in, _ := os.Open("495.in")
	defer in.Close()
	out, _ := os.Create("495.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintf(out, "The Fibonacci number for %d is %v\n", n, &f[n])
	}
}

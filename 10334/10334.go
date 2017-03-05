// UVa 10334 - Ray Through Glasses

package main

import (
	"fmt"
	"math/big"
	"os"
)

const max = 1001

var f = func() [max]big.Int {
	var f [max]big.Int
	f[0].SetInt64(1)
	f[1].SetInt64(2)
	for i := 2; i < max; i++ {
		f[i].Add(&f[i-2], &f[i-1])
	}
	return f
}()

func main() {
	in, _ := os.Open("10334.in")
	defer in.Close()
	out, _ := os.Create("10334.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, &f[n])
	}
}

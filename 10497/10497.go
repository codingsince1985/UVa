// UVa 10497 - Sweet Child Makes Trouble

package main

import (
	"fmt"
	"math/big"
	"os"
)

const max = 800

var f = func() []big.Int {
	f := make([]big.Int, max+1)
	f[0].SetInt64(0)
	f[1].SetInt64(0)
	f[2].SetInt64(1)
	for i := 3; i <= max; i++ {
		f[i].Add(&f[i-1], &f[i-2]).Mul(&f[i], big.NewInt(int64(i-1)))
	}
	return f
}()

func main() {
	in, _ := os.Open("10497.in")
	defer in.Close()
	out, _ := os.Create("10497.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == -1 {
			break
		}
		fmt.Fprintln(out, &f[n])
	}
}

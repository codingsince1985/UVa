// UVa 10375 - Choose and divide

package main

import (
	"fmt"
	"math/big"
	"os"
)

const max = 10000

var f = func() []big.Float {
	f := make([]big.Float, max)
	for i, p := 1, big.NewFloat(1); i < max; i++ {
		f[i].Set(p.Mul(p, big.NewFloat(float64(i))))
	}
	return f
}()

func solve(p, q, r, s int) float64 {
	var ans big.Float
	v, _ := ans.Mul(&f[p], &f[r-s]).Mul(&ans, &f[s]).Quo(&ans, &f[r]).Quo(&ans, &f[p-q]).Quo(&ans, &f[q]).Float64()
	return v
}

func main() {
	in, _ := os.Open("10375.in")
	defer in.Close()
	out, _ := os.Create("10375.out")
	defer out.Close()

	var p, q, r, s int
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d%d", &p, &q, &r, &s); err != nil {
			break
		}
		fmt.Fprintf(out, "%.5f\n", solve(p, q, r, s))
	}
}

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
	p := big.NewFloat(1)
	for i := 1; i < max; i++ {
		p.Mul(p, big.NewFloat(float64(i)))
		f[i].Set(p)
	}
	return f
}()

func solve(p, q, r, s int) float64 {
	var ans big.Float
	ans.Mul(&f[p], &f[r-s])
	ans.Mul(&ans, &f[s])
	ans.Quo(&ans, &f[r])
	ans.Quo(&ans, &f[p-q])
	ans.Quo(&ans, &f[q])
	v, _ := ans.Float64()
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

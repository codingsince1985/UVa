// UVa 10925 - Krakovia

package main

import (
	"fmt"
	"math/big"
	"os"
)

func solve(f int, costs []big.Int) (*big.Int, *big.Int) {
	var sum, bill big.Int
	for _, cost := range costs {
		sum.Add(&sum, &cost)
	}
	return &sum, bill.Set(&sum).Div(&bill, big.NewInt(int64(f)))
}

func main() {
	in, _ := os.Open("10925.in")
	defer in.Close()
	out, _ := os.Create("10925.out")
	defer out.Close()

	var n, f int
	var cost string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &n, &f); n == 0 && f == 0 {
			break
		}
		costs := make([]big.Int, n)
		for i := range costs {
			fmt.Fscanf(in, "%s", &cost)
			costs[i].SetString(cost, 10)
		}
		sum, bill := solve(f, costs)
		fmt.Fprintf(out, "Bill #%d costs %s: each friend should pay %s\n\n", kase, sum, bill)
	}
}

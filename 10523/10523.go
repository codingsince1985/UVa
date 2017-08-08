// UVa 10523 - Very Easy !!!

package main

import (
	"fmt"
	"math/big"
	"os"
)

func solve(n, a int64) *big.Int {
	var tmp big.Int
	part, total, self := big.NewInt(1), big.NewInt(0), big.NewInt(a)
	for i := int64(1); i <= n; i++ {
		part.Mul(part, self)
		tmp.Mul(part, big.NewInt(i))
		total.Add(total, &tmp)
	}
	return total
}

func main() {
	in, _ := os.Open("10523.in")
	defer in.Close()
	out, _ := os.Create("10523.out")
	defer out.Close()

	var n, a int64
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &a); err != nil {
			break
		}
		fmt.Fprintf(out, "%v\n", solve(n, a))
	}
}

// UVa 10106 - Product

package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	in, _ := os.Open("10106.in")
	defer in.Close()
	out, _ := os.Create("10106.out")
	defer out.Close()

	var a, b big.Int
	var c, d string
	for {
		if _, err := fmt.Fscanf(in, "%s\n%s", &c, &d); err != nil {
			break
		}
		a.SetString(c, 10)
		b.SetString(d, 10)
		a.Mul(&a, &b)
		fmt.Fprintf(out, "%v\n", &a)
	}
}

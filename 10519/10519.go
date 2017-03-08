// UVa 10519 - !! Really Strange !!

package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	in, _ := os.Open("10519.in")
	defer in.Close()
	out, _ := os.Create("10519.out")
	defer out.Close()

	var s string
	var n, tmp big.Int
	two := big.NewInt(2)
	for {
		if _, err := fmt.Fscanf(in, "%s", &s); err != nil {
			break
		}
		if s == "0" {
			fmt.Fprintln(out, 1)
		} else {
			n.SetString(s, 10)
			tmp.Mul(&n, &n)
			tmp.Sub(&tmp, &n)
			tmp.Add(&tmp, two)
			fmt.Fprintln(out, &tmp)
		}
	}
}

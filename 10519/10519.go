// UVa 10519 - !! Really Strange !!

package main

import (
	"fmt"
	"math/big"
	"os"
)

func main() {
	var s string
	var n, two, tmp big.Int
	two.SetInt64(2)

	in, _ := os.Open("10519.in")
	out, _ := os.Create("10519.out")
	for {
		_, err := fmt.Fscanf(in, "%s", &s)
		if err != nil {
			break
		}
		if s == "0" {
			fmt.Fprintln(out, 1)
			continue
		}
		n.SetString(s, 10)
		tmp.Mul(&n, &n)
		tmp.Sub(&tmp, &n)
		tmp.Add(&tmp, &two)
		fmt.Fprintln(out, &tmp)
	}
	in.Close()
	out.Close()
}

// UVa 113 - Power of Cryptography

// k^n = p
// ln(k) = ln(p) / n
// k = e^(ln(p)/n)

package main

import (
	"fmt"
	"math/big"
	"os"
)

const max_k = 1000000000

func main() {
	in, _ := os.Open("113.in")
	defer in.Close()
	out, _ := os.Create("113.out")
	defer out.Close()

	var ps string
	var p, tmp big.Int
	var ns int64
	for {
		if _, err := fmt.Fscanf(in, "%d\n%s", &ns, &ps); err != nil {
			break
		}
		p.SetString(ps, 10)
		n := big.NewInt(ns)
		if p.Cmp(n) == 0 {
			fmt.Fprintln(out, 1)
		} else {
			var i1, i2 int64 = 0, max_k
			for {
				var i int64 = (i1 + i2 + 1) / 2
				k := big.NewInt(i)
				tmp.Exp(k, n, nil)
				if tmp.Cmp(&p) == 0 {
					fmt.Fprintln(out, i)
					break
				}
				if tmp.Cmp(&p) > 0 {
					i2 = i
				} else {
					i1 = i
				}
			}
		}
	}
}

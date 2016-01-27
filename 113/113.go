// UVa 113 - Power of Cryptography

// k^n = p
// ln(k) = ln(p) / n
// k = e^(ln(p)/n)

package main

import (
	"fmt"
	"os"
	"math/big"
)

const MAX_K = 1000000000

func main() {
	in, _ := os.Open("113.in")
	defer in.Close()
	out, _ := os.Create("113.out")
	defer out.Close()

	var ns int
	var ps string
	var p, k, n, tmp big.Int
	var i, i1, i2 int64
	for {
		if _, err := fmt.Fscanf(in, "%d\n%s", &ns, &ps); err != nil {
			break
		}
		p.SetString(ps, 10)
		n.SetInt64(int64(ns))

		if p.Cmp(&n) == 0 {
			fmt.Fprintln(out, 1)
			continue
		}

		i1, i2 = 0, MAX_K
		for {
			i = (i1 + i2 + 1) / 2
			k.SetInt64(i)
			tmp.Exp(&k, &n, nil)
			if tmp.Cmp(&p) == 0 {
				fmt.Fprintln(out, i)
				break
			} else if tmp.Cmp(&p) > 0 {
				i2 = i
			} else {
				i1 = i
			}
		}
	}
}

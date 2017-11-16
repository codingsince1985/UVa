// UVa 10023 - Square root

package main

import (
	"fmt"
	"math/big"
	"os"
)

var two = big.NewInt(2)

func sqrt(num string) *big.Int {
	var n, m, h, tmp big.Int
	n.SetString(num, 10)
	h.Set(&n)
	l := big.NewInt(1)
	for {
		m.Add(l, &h).Div(&m, two)
		cmp := tmp.Mul(&m, &m).Cmp(&n)
		switch {
		case cmp == 0:
			return &m
		case cmp < 0:
			l.Set(&m)
		default:
			h.Set(&m)
		}
	}
}

func main() {
	in, _ := os.Open("10023.in")
	defer in.Close()
	out, _ := os.Create("10023.out")
	defer out.Close()

	var kase int
	var line string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%s", &line)
		fmt.Fprintln(out, sqrt(line))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}

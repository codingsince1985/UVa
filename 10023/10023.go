// UVa 10023 - Square root

package main

import (
	"fmt"
	"math/big"
	"os"
)

var two = big.NewInt(2)

func sqrt(num string) *big.Int {
	var n, l, m, h, tmp big.Int
	n.SetString(num, 10)
	l.SetInt64(1)
	h.Set(&n)
	for {
		m.Add(&l, &h)
		m.Div(&m, two)
		tmp.Mul(&m, &m)
		cmp := tmp.Cmp(&n)
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
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%s", &line)
		fmt.Fprintf(out, "%v\n", sqrt(line))
		if kase--; kase > 0 {
			fmt.Fprintln(out)
		}
	}
}

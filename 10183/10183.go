// UVa 10183 - How Many Fibs?

package main

import (
	"fmt"
	"math/big"
	"os"
)

const MAX = 101

var p []*big.Int

func prepare() {
	p = make([]*big.Int, 0)
	p0 := big.NewInt(1)
	p = append(p, p0)
	p1 := big.NewInt(1)
	p = append(p, p1)

	var tmp *big.Int
	var s string
	for i := 2;; i ++ {
		tmp = big.NewInt(0)
		tmp = tmp.Add(p[i - 2], p[i - 1])
		s = fmt.Sprintf("%v", tmp)
		if len(s) > MAX {
			break;
		}
		p = append(p, tmp)
	}
}

func do() {
	in, _ := os.Open("10183.in")
	out, _ := os.Create("10183.out")

	var a, b string
	var c int
	n1 := big.NewInt(0)
	n2 := big.NewInt(0)
	l := len(p)
	for {
		fmt.Fscanf(in, "%s %s", &a, &b)
		if a == "0" && b == "0" {
			break;
		}
		n1.SetString(a, 10)
		n2.SetString(b, 10)

		c = 0
		for i := 1; i < l; i ++ {
			if p[i].Cmp(n1) >= 0 && p[i].Cmp(n2) <= 0 {
				c++
			}
			if p[i].Cmp(n2) > 0 {
				break;
			}
		}
		fmt.Fprintln(out, c)
	}

	in.Close()
	out.Close()
}

func main() {
	prepare()
	do()
}
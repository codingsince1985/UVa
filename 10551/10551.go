// UVa 10551 - Basic Remains

package main

import (
	"fmt"
	"math/big"
	"os"
)

func solve(base int, s1, s2 string) string {
	var p, m big.Int
	p.SetString(s1, base)
	m.SetString(s2, base)
	return p.Mod(&p, &m).Text(base)
}

func main() {
	in, _ := os.Open("10551.in")
	defer in.Close()
	out, _ := os.Create("10551.out")
	defer out.Close()

	var base int
	var p, m string
	for {
		if fmt.Fscanf(in, "%d", &base); base == 0 {
			break
		}
		fmt.Fscanf(in, "%s%s", &p, &m)
		fmt.Fprintln(out, solve(base, p, m))
	}
}

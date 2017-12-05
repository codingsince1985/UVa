// UVa 10814 - Simplifying Fractions

package main

import (
	"fmt"
	"math/big"
	"os"
)

var zero = big.NewInt(0)

func gcd(a, b *big.Int) *big.Int {
	if b.Cmp(zero) == 0 {
		return a
	}
	return gcd(b, a.Mod(a, b))
}

func solve(n1, n2 string) (string, string) {
	var a, b, d big.Int
	a.SetString(n1, 10)
	b.SetString(n2, 10)
	d.Set(gcd(&a, &b))

	a.SetString(n1, 10)
	b.SetString(n2, 10)
	return a.Div(&a, &d).String(), b.Div(&b, &d).String()
}

func main() {
	in, _ := os.Open("10814.in")
	defer in.Close()
	out, _ := os.Create("10814.out")
	defer out.Close()

	var kase int
	var n1, n2 string
	var ch byte
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		if _, err := fmt.Fscanf(in, "%s %c %s", &n1, &ch, &n2); err != nil {
			break
		}
		a, b := solve(n1, n2)
		fmt.Fprintf(out, "%s / %s\n", a, b)
	}
}

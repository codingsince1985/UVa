// UVa 10006 - Carmichael Numbers

package main

import (
	"fmt"
	"math/big"
	"os"
)

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isCarmichael(n int) bool {
	var b, e, p big.Int
	e.SetInt64(int64(n))
	for a := 2; a < n; a++ {
		b.SetInt64(int64(a))
		p.Exp(&b, &e, &e)
		p.Sub(&p, &b)
		if p.Int64() != 0 {
			return false
		}
	}
	return !isPrime(n)
}

func main() {
	in, _ := os.Open("10006.in")
	defer in.Close()
	out, _ := os.Create("10006.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		if isCarmichael(n) {
			fmt.Fprintf(out, "The number %d is a Carmichael number.\n", n)
		} else {
			fmt.Fprintf(out, "%d is normal.\n", n)
		}
	}
}

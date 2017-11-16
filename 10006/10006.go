// UVa 10006 - Carmichael Numbers

package main

import (
	"fmt"
	"math/big"
	"os"
)

func isPrime(n int64) bool {
	if n%2 == 0 {
		return false
	}
	for i := int64(3); i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isCarmichael(n int64) bool {
	var b, p big.Int
	e := big.NewInt(n)
	for a := int64(2); a < n; a++ {
		b.SetInt64(a)
		if p.Exp(&b, e, e).Sub(&p, &b); p.Int64() != 0 {
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

	var n int64
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

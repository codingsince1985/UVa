// UVa 10852 - Less Prime

package main

import (
	"fmt"
	"os"
)

const max = 10000

var primes = sieve()

func sieve() []bool {
	p := make([]bool, max+1)
	for i := 2; i*i <= max; i++ {
		if !p[i] {
			for j := 2 * i; j <= max; j += i {
				p[j] = true
			}
		}
	}
	return p
}

func solve(n int) int {
	for i := n/2 + 1; i <= n; i++ {
		if !primes[i] {
			return i
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("10852.in")
	defer in.Close()
	out, _ := os.Create("10852.out")
	defer out.Close()

	var kase, n int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d", &n)
		fmt.Fprintln(out, solve(n))
	}
}

// UVa 10780 - Again Prime? No Time.

package main

import (
	"fmt"
	"math"
	"os"
)

const max = 10000

var primes = sieve()

func sieve() []bool {
	p := make([]bool, max+1)
	p[0], p[1] = true, true
	for i := 2; i*i <= max; i++ {
		if !p[i] {
			for j := 2 * i; j <= max; j += i {
				p[j] = true
			}
		}
	}
	return p
}

func primeFactorize(n int) map[int]int {
	factorMap := make(map[int]int)
	for p := 2; p <= n && n != 1; p++ {
		if !primes[p] {
			for ; n%p == 0; n /= p {
				factorMap[p]++
			}
		}
	}
	return factorMap
}

func solve(m, n int) int {
	factorMap := make(map[int]int)
	for i := 2; i <= n; i++ {
		for k, v := range primeFactorize(i) {
			factorMap[k] += v
		}
	}
	min := math.MaxInt32
	for k, v := range primeFactorize(m) {
		if f := factorMap[k] / v; f < min {
			min = f
		}
	}
	return min
}

func main() {
	in, _ := os.Open("10780.in")
	defer in.Close()
	out, _ := os.Create("10780.out")
	defer out.Close()

	var i, m, n int
	fmt.Fscanf(in, "%d", &i)
	for kase := 1; kase <= i; kase++ {
		fmt.Fscanf(in, "%d%d", &m, &n)
		fmt.Fprintf(out, "Case %d:\n", kase)
		if power := solve(m, n); power == math.MaxInt32 || power == 0 {
			fmt.Fprintln(out, "Impossible to divide")
		} else {
			fmt.Fprintln(out, power)
		}
	}
}

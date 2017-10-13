// UVa 10738 - Riemann vs Mertens

package main

import (
	"fmt"
	"os"
)

const max = 1000000

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
	factors := make(map[int]int)
	for i, v := range primes {
		if !v {
			for ; n%i == 0; n /= i {
				factors[i]++
			}
			if n == 1 {
				break
			}
		}
	}
	return factors
}

func mobius(n int) int {
	factors := primeFactorize(n)
	for _, v := range factors {
		if v > 1 {
			return 0
		}
	}
	if len(factors)%2 == 0 {
		return 1
	}
	return -1
}

func solve(n int, mu map[int]int) (int, int) {
	var sum int
	for i := 1; i <= n; i++ {
		if _, ok := mu[i]; !ok {
			mu[i] = mobius(i)
		}
		sum += mu[i]
	}
	return mu[n], sum
}

func main() {
	in, _ := os.Open("10738.in")
	defer in.Close()
	out, _ := os.Create("10738.out")
	defer out.Close()

	var n int
	mu := map[int]int{1: 1}
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		mun, mn := solve(n, mu)
		fmt.Fprintf(out, "%8d%8d%8d\n", n, mun, mn)
	}
}

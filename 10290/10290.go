// UVa 10290 - {Sum+=i++} to Reach N

package main

import (
	"fmt"
	"os"
)

const max = 3000000

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

func primeFactorize(n int64) map[int64]int {
	factorMap := make(map[int64]int)
	for p := int64(2); p <= n && n != 1; p++ {
		if !primes[p] {
			for ; n%p == 0; n /= p {
				factorMap[p]++
			}
		}
	}
	if n != 1 {
		factorMap[n] = 1
	}
	return factorMap
}

func solve(n int64) int {
	if n == 0 {
		return 1
	}
	count := 1
	for k, v := range primeFactorize(n) {
		if k != 2 {
			count *= v + 1
		}
	}
	return count
}

func main() {
	in, _ := os.Open("10290.in")
	defer in.Close()
	out, _ := os.Create("10290.out")
	defer out.Close()

	var n int64
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, solve(n))
	}
}

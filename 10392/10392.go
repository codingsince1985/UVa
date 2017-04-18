// UVa 10392 - Factoring Large Numbers

package main

import (
	"fmt"
	"os"
)

const max = 10000000

var (
	out    *os.File
	primes = sieve()
)

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

func isPrime(n int) bool {
	if n <= max {
		return !primes[n]
	}
	for i := range primes {
		if i*i > n {
			break
		}
		if !primes[i] && n%i == 0 {
			return false
		}
	}
	return true
}

func primeFactorize(n int64) {
	for idx := 2; n > 1 && int64(idx)*int64(idx) <= n; idx++ {
		if isPrime(idx) {
			for n%int64(idx) == 0 {
				n /= int64(idx)
				fmt.Fprintf(out, "    %d\n", idx)
			}
		}
	}
	if n > 1 {
		fmt.Fprintf(out, "    %d\n", n)
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("10392.in")
	defer in.Close()
	out, _ = os.Create("10392.out")
	defer out.Close()

	var n int64
	for {
		if fmt.Fscanf(in, "%d", &n); n < 0 {
			break
		}
		primeFactorize(n)
	}
}

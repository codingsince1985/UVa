// UVa 10539 - Almost Prime Numbers

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

func isPrime(n int64) bool {
	if n <= max {
		return !primes[n]
	}
	for i := range primes {
		if !primes[i] && int64(i*i) <= n && n%int64(i) == 0 {
			return false
		}
	}
	return true
}

func solve(low, high int64) int {
	var i, almost int64
	var cnt int
	for i = 2; i*i <= high; i++ {
		if i*i >= low && isPrime(i) {
			for almost = i * i; almost <= high; almost = almost * i {
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	in, _ := os.Open("10539.in")
	defer in.Close()
	out, _ := os.Create("10539.out")
	defer out.Close()

	var n, low, high int64
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d", &low, &high)
		fmt.Fprintln(out, solve(low, high))
	}
}

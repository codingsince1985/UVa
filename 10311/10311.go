// UVa 10311 - Goldbach and Euler

package main

import (
	"fmt"
	"os"
)

const max = 10000000

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

func isPrime(n int) bool {
	if n <= 0 {
		return false
	}
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

func solve(n int) (int, int) {
	if n%2 == 0 {
		for i := n / 2; i >= 3; i-- {
			if isPrime(i) && isPrime(n-i) {
				return i, n - i
			}
		}
		return -1, -1
	}
	if isPrime(n - 2) {
		return 2, n - 2
	}
	return -1, -1
}

func main() {
	in, _ := os.Open("10311.in")
	defer in.Close()
	out, _ := os.Create("10311.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		if n1, n2 := solve(n); n1 == -1 {
			fmt.Fprintf(out, "%d is not the sum of two primes!\n", n)
		} else {
			fmt.Fprintf(out, "%d is the sum of %d and %d.\n", n, n1, n2)
		}
	}
}

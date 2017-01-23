// UVa 10699 - Count the factors

package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
	if n%2 == 0 {
		return n == 2
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primeList() []int {
	var primes []int
	primes = append(primes, 2)
	for i := 3; i < 100000; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func factorize(n int, primes []int) map[int]bool {
	factors := make(map[int]bool)
	for _, v := range primes {
		for n%v == 0 {
			n /= v
			factors[v] = true
		}
		if isPrime(n) {
			factors[n] = true
			break
		}
	}
	return factors
}

func main() {
	in, _ := os.Open("10699.in")
	defer in.Close()
	out, _ := os.Create("10699.out")
	defer out.Close()

	primes := primeList()
	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintf(out, "%d : %d\n", n, len(factorize(n, primes)))
	}
}

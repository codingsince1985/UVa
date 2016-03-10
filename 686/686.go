// Uva 686 - Goldbach's Conjecture (II)

package main

import (
	"fmt"
	"os"
)

func isPrime(n int) bool {
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

var primes = func() map[int]bool {
	primes := make(map[int]bool)
	primes[2] = true
	for i := 3; i < 32768; i += 2 {
		if isPrime(i) {
			primes[i] = true
		}
	}
	return primes
}()

func solve(n int) int {
	if n == 4 {
		return 1
	}
	cnt := 0
	for i := 3; i <= n/2; i += 2 {
		if primes[i] && primes[n-i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	in, _ := os.Open("686.in")
	defer in.Close()
	out, _ := os.Create("686.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintln(out, solve(n))
	}
}

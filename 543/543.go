// UVa 543 - Goldbach's Conjecture

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

var primeSet = func() map[int]bool {
	primeSet := make(map[int]bool)
	for i := 3; i < 1000000; i += 2 {
		if isPrime(i) {
			primeSet[i] = true
		}
	}
	return primeSet
}()

func solve(out *os.File, n int) {
	for i := 3; i <= n/2; i += 2 {
		if primeSet[i] && primeSet[n-i] {
			fmt.Fprintf(out, "%d = %d + %d\n", n, i, n-i)
			return
		}
	}
	fmt.Fprintln(out, "Goldbach's conjecture is wrong.")
}

func main() {
	in, _ := os.Open("543.in")
	defer in.Close()
	out, _ := os.Create("543.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		solve(out, n)
	}
}

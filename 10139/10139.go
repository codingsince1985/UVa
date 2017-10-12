// UVa 10139 - Factovisors

package main

import (
	"fmt"
	"os"
)

var (
	pl    = make(map[int]bool)
	cache = make(map[int]map[int]int)
)

func isPrime(n int) bool {
	if res, ok := pl[n]; ok {
		return res
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			pl[n] = false
			return false
		}
	}
	pl[n] = true
	return true
}

func factorize(n int) map[int]int {
	f := make(map[int]int)
	for i := 2; n > 1; i++ {
		if isPrime(i) {
			for n != 1 && n%i == 0 {
				f[i]++
				n /= i
			}
		}
	}
	return f
}

func primeFactorizeFactorial(n int) map[int]int {
	var ok bool
	facts := make(map[int]int)
	for i := 2; i <= n; i++ {
		if _, ok = cache[i]; !ok {
			cache[i] = factorize(i)
		}
		for k, v := range cache[i] {
			facts[k] += v
		}
	}
	return facts
}

func contains(f1, f2 map[int]int) bool {
	for k, v1 := range f1 {
		if v2, ok := f2[k]; !ok || v2 < v1 {
			return false
		}
	}
	return true
}

func main() {
	in, _ := os.Open("10139.in")
	defer in.Close()
	out, _ := os.Create("10139.out")
	defer out.Close()

	pl[2] = true
	var n, m int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &m); err != nil {
			break
		}
		if contains(factorize(m), primeFactorizeFactorial(n)) {
			fmt.Fprintf(out, "%d divides %d!\n", m, n)
		} else {
			fmt.Fprintf(out, "%d does not divide %d!\n", m, n)
		}
	}
}

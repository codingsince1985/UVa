// UVa 10484 - Divisibility of Factors

package main

import (
	"fmt"
	"os"
)

const max = 100

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

func subtract(f1, f2 map[int]int) map[int]int {
	f3 := make(map[int]int)
	for k, v := range f2 {
		switch {
		case f1[k] == v:
			delete(f1, k)
		case f1[k] > v:
			f1[k] -= v
		default:
			delete(f1, k)
			f3[k] = v - f1[k]
		}
	}
	return f3
}

func solve(n, d int) int {
	factors := make(map[int]int)
	for i := 2; i <= n; i++ {
		for k, v := range primeFactorize(i) {
			factors[k] += v
		}
	}
	if len(subtract(factors, primeFactorize(d))) > 0 {
		return 0
	}
	count := 1
	for _, v := range factors {
		count *= v + 1
	}
	return count
}

func main() {
	in, _ := os.Open("10484.in")
	defer in.Close()
	out, _ := os.Create("10484.out")
	defer out.Close()

	var n, d int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &d); n == 0 && d == 0 {
			break
		}
		fmt.Fprintln(out, solve(n, d))
	}
}

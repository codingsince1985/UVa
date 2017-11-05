// UVa 11466 - Largest Prime Divisor

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

func solve(n int64) int64 {
	if n < 0 {
		n = -n
	}
	var count, largest int64
	for i := int64(0); i <= int64(max) && i <= n; i++ {
		if !primes[i] {
			for ; n%i == 0; n /= i {
				if largest != i {
					count++
					largest = i
				}
			}
		}
	}
	if n != 1 {
		count++
		largest = n
	}
	if count == 1 {
		return -1
	}
	return largest
}

func main() {
	in, _ := os.Open("11466.in")
	defer in.Close()
	out, _ := os.Create("11466.out")
	defer out.Close()

	var n int64
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintln(out, solve(n))
	}
}

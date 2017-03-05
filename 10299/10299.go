// UVa 10299 - Relatives

package main

import (
	"fmt"
	"os"
)

const max = 31622

var primes = func() []int {
	p := make([]bool, max+1)
	p[0], p[1] = true, true
	for i := 2; i*i <= max; i++ {
		if !p[i] {
			for j := 2 * i; j <= max; j += i {
				p[j] = true
			}
		}
	}
	var ps []int
	for i, vi := range p {
		if !vi {
			ps = append(ps, i)
		}
	}
	return ps
}()

func solve(n int) int {
	if n == 1 {
		return 0
	}
	result := n
	for _, prime := range primes {
		if n < prime {
			break
		}
		isFactor := false
		for n%prime == 0 {
			n /= prime
			isFactor = true
		}
		if isFactor {
			result = result / prime * (prime - 1)
		}
	}
	if n != 1 {
		result = result / n * (n - 1)
	}
	return result
}

func main() {
	in, _ := os.Open("10299.in")
	defer in.Close()
	out, _ := os.Create("10299.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintln(out, solve(n))
	}
}

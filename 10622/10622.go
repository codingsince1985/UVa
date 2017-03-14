// UVa 10622 - Perfect P-th Powers

package main

import (
	"fmt"
	"os"
)

const max = 46340

var primes = sieve()

func sieve() []int {
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
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func primeFactor(n int) map[int]int {
	pf := make(map[int]int)
	for n > 1 {
		for _, prime := range primes {
			for n%prime == 0 {
				pf[prime]++
				n /= prime
			}
			if n == 1 {
				break
			}
		}
	}
	return pf
}

func main() {
	in, _ := os.Open("10622.in")
	defer in.Close()
	out, _ := os.Create("10622.out")
	defer out.Close()

	var x int
	var negative bool
	for {
		if fmt.Fscanf(in, "%d", &x); x == 0 {
			break
		}
		if x < 0 {
			negative = true
			x = -x
		}
		pf := primeFactor(x)
		g := 0
		for _, v := range pf {
			if g == 0 {
				g = v
			} else {
				g = gcd(g, v)
			}
		}
		if negative && g != 0 {
			for g%2 == 0 {
				g /= 2
			}
		}
		fmt.Fprintln(out, g)
	}
}

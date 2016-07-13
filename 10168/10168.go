// UVa 10168 - Summation of Four Primes

package main

import (
	"fmt"
	"os"
)

const MAX = 10000000

func sieve() []bool {
	p := make([]bool, MAX+1)
	for i := 2; i*i <= MAX; i++ {
		if !p[i] {
			for j := 2 * i; j <= MAX; j += i {
				p[j] = true
			}
		}
	}
	return p
}

func main() {
	in, _ := os.Open("10168.in")
	defer in.Close()
	out, _ := os.Create("10168.out")
	defer out.Close()

	var n int
	p := sieve()
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		if n < 8 {
			fmt.Fprintln(out, "Impossible.")
			continue
		}
		switch {
		case n%2 == 0:
			fmt.Fprint(out, "2 2 ")
			n -= 4
		default:
			fmt.Fprint(out, "2 3 ")
			n -= 5
		}
		for i := 2; i < n; i++ {
			if !p[i] && !p[n-i] {
				fmt.Fprintln(out, i, n-i)
				break
			}
		}
	}
}

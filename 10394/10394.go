// UVa 10394 - Twin Primes

package main

import (
	"fmt"
	"os"
)

const max = 20000000

func sieve() []bool {
	p := make([]bool, max+1)
	for i := 2; i*i <= max; i++ {
		if !p[i] {
			for j := i + i; j <= max; j += i {
				p[j] = true
			}
		}
	}
	return p
}

func main() {
	in, _ := os.Open("10394.in")
	defer in.Close()
	out, _ := os.Create("10394.out")
	defer out.Close()

	p := sieve()
	var tp []int
	for i := 3; i+2 <= max; i += 2 {
		if !p[i] && !p[i+2] {
			tp = append(tp, i)
		}
	}

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintf(out, "(%d, %d)\n", tp[n-1], tp[n-1]+2)
	}
}

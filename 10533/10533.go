// UVa 10533 - Digit Primes

package main

import (
	"fmt"
	"os"
)

const max = 1000000

func sieve() []bool {
	p := make([]bool, max+1)
	p[0], p[1] = true, true
	for i := 2; i*i <= max; i++ {
		if !p[i] {
			for j := i + i; j <= max; j += i {
				p[j] = true
			}
		}
	}
	return p
}

func digitSum(n int) int {
	total := 0
	for n > 0 {
		total += n % 10
		n /= 10
	}
	return total
}

func main() {
	in, _ := os.Open("10533.in")
	defer in.Close()
	out, _ := os.Create("10533.out")
	defer out.Close()

	p := sieve()
	var n, t1, t2 int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d", &t1, &t2)
		cnt := 0
		for i := t1; i <= t2; i++ {
			if !p[i] && !p[digitSum(i)] {
				cnt++
			}
		}
		fmt.Fprintln(out, cnt)
	}
}

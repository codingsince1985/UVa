// UVa 10948 - The primary problem

package main

import (
	"fmt"
	"os"
)

const MAX = 1000000

var primes = func() []bool {
	p := make([]bool, MAX+1)
	p[0], p[1] = true, true
	for i := 2; i*i <= MAX; i++ {
		if !p[i] {
			for j := 2 * i; j <= MAX; j += i {
				p[j] = true
			}
		}
	}
	return p
}()

func solve(n int) (int, int) {
	if !primes[n-2] {
		return 2, n - 2
	}
	for i := 3; i <= n/2; i += 2 {
		if !primes[i] && !primes[n-i] {
			return i, n - i
		}
	}
	return -1, -1
}

func main() {
	in, _ := os.Open("10948.in")
	defer in.Close()
	out, _ := os.Create("10948.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintf(out, "%d:\n", n)
		if n1, n2 := solve(n); n1 == -1 && n2 == -1 {
			fmt.Fprintln(out, "NO WAY!")
		} else {
			fmt.Fprintf(out, "%d+%d\n", n1, n2)
		}
	}
}

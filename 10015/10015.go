// UVa 10015 - Joseph's Cousin

package main

import (
	"fmt"
	"os"
)

const MAX = 3500

func isPrime(n int) bool {
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primes() []int {
	p := make([]int, MAX)
	p[0] = 2
	for i, idx := 3, 1; idx < MAX; i += 2 {
		if isPrime(i) {
			p[idx] = i
			idx++
		}
	}
	return p
}

func main() {
	in, _ := os.Open("10015.in")
	defer in.Close()
	out, _ := os.Create("10015.out")
	defer out.Close()

	var n int
	p := primes()
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		left, idx := n, -1
		killed := make([]bool, n)
		for left > 1 {
			cnt := p[n-left]
			if cnt > left {
				cnt = cnt % left
			}
			for cnt > 0 {
				idx++
				if idx == n {
					idx = 0
				}
				if !killed[idx] {
					cnt--
				}
			}
			killed[idx] = true
			left--
		}
		for i, v := range killed {
			if !v {
				fmt.Fprintln(out, i+1)
				break
			}
		}
	}
}

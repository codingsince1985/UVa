// UVa 10140 - Prime Distance

package main

import (
	"fmt"
	"math"
	"os"
)

const MAX = 46340 // √(2^31-1)

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

func isPrime(n int) bool {
	if n <= MAX {
		return !primes[n]
	}
	for i := range primes {
		if !primes[i] && n%i == 0 {
			return false
		}
	}
	return true
}

func solve(l, u int) (c1, c2, d1, d2 int) {
	for i, pre, min, max := l, -1, math.MaxInt32, math.MinInt32; i <= u; i++ {
		if isPrime(i) {
			if pre != -1 {
				distance := i - pre
				if distance < min {
					min, c1, c2 = distance, pre, i
				}
				if distance > max {
					max, d1, d2 = distance, pre, i
				}
			}
			pre = i
		}
	}
	return
}

func main() {
	in, _ := os.Open("10140.in")
	defer in.Close()
	out, _ := os.Create("10140.out")
	defer out.Close()

	var l, u int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &l, &u); err != nil {
			break
		}
		if c1, c2, d1, d2 := solve(l, u); c1 == 0 && c2 == 0 && d1 == 0 && d2 == 0 {
			fmt.Fprintln(out, "There are no adjacent primes.")
		} else {
			fmt.Fprintf(out, "%d,%d are closest, %d,%d are most distant.\n", c1, c2, d1, d2)
		}
	}
}

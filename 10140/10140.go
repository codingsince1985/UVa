// UVa 10140 - Prime Distance

package main

import (
	"fmt"
	"math"
	"os"
)

const max = 46340 // √(2^31-1)

var primes = func() []bool {
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
}()

func isPrime(n int) bool {
	if n <= max {
		return !primes[n]
	}
	for i := range primes {
		if i*i > n {
			break
		}
		if !primes[i] && n%i == 0 {
			return false
		}
	}
	return true
}

func solve(l, u int) (c1, c2, d1, d2 int) {
	pre, min, max := -1, math.MaxInt32, math.MinInt32
	for i := l; i <= u; i++ {
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

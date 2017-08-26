// UVa 10791 - Minimum Sum LCM

package main

import (
	"fmt"
	"math"
	"os"
)

const max = 46341 // √(2^31-1)+1

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
	var primeCount, ans int64
	tmp := n
	sqrt := int64(math.Sqrt(float64(n)) + .5)
	for i := int64(2); i <= sqrt && tmp != 0 && i <= tmp; i++ {
		if !primes[i] && tmp%i == 0 {
			primeCount++
			primeFactor := int64(1)
			for ; tmp%i == 0; tmp /= i {
				primeFactor *= i
			}
			ans += primeFactor
		}
	}
	switch {
	case tmp == n:
		return n + 1
	case primeCount == 1 || tmp != 1:
		return ans + tmp
	default:
		return ans
	}
}

func main() {
	in, _ := os.Open("10791.in")
	defer in.Close()
	out, _ := os.Create("10791.out")
	defer out.Close()

	var n int64
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintf(out, "Case %d: %d\n", kase, solve(n))
	}
}

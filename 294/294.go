// UVa 294 - Divisors

// If a number N == a^i * b^j * ... * c^k, then N has (i+1)*(j+1)*...*(k+1) divisors.
// If d is a divisor of n, then so is n/d, but d & n/d cannot be both greater than sqrt(n).

package main

import (
	"fmt"
	"os"
	"math"
)

var p map[int]bool

func factorize(n int) map[int]int {
	f := make(map[int]int)
	if n == 1 {
		return nil
	}

	t := 2
	done := false
	for !done {
		for n % t == 0 {
			f[t] ++
			if n /= t; n == 1 {
				done = true
				break
			}
		}

		if !done {
			t ++
			if t > int(math.Sqrt(float64(n)) + .5) {
				// remaining must be a prime
				f[n] = 1
				break
			}
		}
	}
	return f
}

func numberOfDivisors(f map[int]int) int {
	p := 1
	if f != nil {
		for _, value := range f {
			p *= (value + 1)
		}
	}
	return p
}

func main() {
	p = make(map[int]bool)
	p[2] = true

	in, _ := os.Open("294.in")
	defer in.Close()
	out, _ := os.Create("294.out")
	defer out.Close()

	var n, L, U int
	fmt.Fscanf(in, "%d", &n)
	for i := 0; i < n; i ++ {
		var num, max, new int
		fmt.Fscanf(in, "%d %d", &L, &U)
		for j := L; j <= U; j ++ {
			factors := factorize(j)
			if new = numberOfDivisors(factors); new > max {
				max = new
				num = j
			}
		}
		fmt.Fprintf(out, "Between %d and %d, %d has a maximum of %d divisors.\n", L, U, num, max)
	}
}

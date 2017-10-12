// UVa 294 - Divisors

// If a number N == a^i * b^j * ... * c^k, then N has (i+1)*(j+1)*...*(k+1) divisors.
// If d is a divisor of n, then so is n/d, but d & n/d cannot be both greater than sqrt(n).

package main

import (
	"fmt"
	"os"
)

func factorize(n int) map[int]int {
	if n == 1 {
		return nil
	}

	f := make(map[int]int)
here:
	for t := 2; t*t <= n; t++ {
		for n%t == 0 {
			f[t]++
			if n /= t; n == 1 {
				break here
			}
		}
	}
	if n != 1 {
		f[n] = 1
	}
	return f
}

func numberOfDivisors(f map[int]int) int {
	p := 1
	if f != nil {
		for _, value := range f {
			p *= value + 1
		}
	}
	return p
}

func main() {
	in, _ := os.Open("294.in")
	defer in.Close()
	out, _ := os.Create("294.out")
	defer out.Close()

	var n, L, U int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		var num, max, divisors int
		fmt.Fscanf(in, "%d%d", &L, &U)
		for j := L; j <= U; j++ {
			factors := factorize(j)
			if divisors = numberOfDivisors(factors); divisors > max {
				max = divisors
				num = j
			}
		}
		fmt.Fprintf(out, "Between %d and %d, %d has a maximum of %d divisors.\n", L, U, num, max)
	}
}

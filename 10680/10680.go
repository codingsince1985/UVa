// UVa 10680 - LCM

package main

import (
	"fmt"
	"os"
)

const (
	max = 1000000
	mod = 1000000000
)

var f = func() []int {
	factors := sieve()
	f := make([]int, max+1)
	f[1] = 1
	for i := 2; i <= max; i++ {
		f[i] = (f[i-1] * factors[i]) % mod
	}
	return f
}()

func sieve() []int {
	factors := make([]int, max+1)
	for i := range factors {
		factors[i] = 1
	}
	p := make([]bool, max+1)
	p[0], p[1] = true, true
	for i := 2; i*i <= max; i++ {
		if !p[i] {
			for j := 2 * i; j <= max; j += i {
				p[j] = true
			}
			for j := i; j <= max; j *= i {
				factors[j] = i
			}
		}
	}
	return factors
}

func main() {
	in, _ := os.Open("10680.in")
	defer in.Close()
	out, _ := os.Create("10680.out")
	defer out.Close()

	var n, d int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		for d = f[n]; d%10 == 0; d /= 10 {
		}
		fmt.Fprintln(out, d%10)
	}
}

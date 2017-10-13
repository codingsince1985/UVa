// UVa 10820 - Send a Table

package main

import (
	"fmt"
	"os"
)

const max = 50000

// Euler's totient function
var phi = func() []int {
	phi := make([]int, max+1)
	for i := 2; i <= max; i++ {
		if phi[i] == 0 {
			for j := i; j <= max; j += i {
				if phi[j] == 0 {
					phi[j] = j
				}
				phi[j] = phi[j] / i * (i - 1)
			}
		}
	}
	return phi
}()

func solve(n int) int {
	if n == 1 {
		return 1
	}
	var sum int
	for i := 2; i <= n; i++ {
		sum += phi[i]
	}
	return 2*sum + 1
}

func main() {
	in, _ := os.Open("10820.in")
	defer in.Close()
	out, _ := os.Create("10820.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintln(out, solve(n))
	}
}

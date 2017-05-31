// UVa 10212 - The Last Non-zero Digit.

package main

import (
	"fmt"
	"os"
)

func solve(n, m uint64) uint64 {
	var product uint64 = 1
	for i := n - m + 1; i <= n; i++ {
		product *= uint64(i)
		for product%10 == 0 {
			product /= 10
		}
		product %= 1000000000
	}
	return product % 10
}

func main() {
	in, _ := os.Open("10212.in")
	defer in.Close()
	out, _ := os.Create("10212.out")
	defer out.Close()

	var n, m uint64
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &m); err != nil {
			break
		}
		fmt.Fprintln(out, solve(n, m))
	}
}

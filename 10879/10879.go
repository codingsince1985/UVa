// UVa 10879 - Code Refactoring

package main

import (
	"fmt"
	"os"
)

func solve(k int) [2][2]int {
	var factors [2][2]int
	for idx, i := 0, 2; idx < 2 && i*i < k; i++ {
		if k%i == 0 {
			factors[idx] = [2]int{i, k / i}
			idx++
		}
	}
	return factors
}

func main() {
	in, _ := os.Open("10879.in")
	defer in.Close()
	out, _ := os.Create("10879.out")
	defer out.Close()

	var n, k int
	fmt.Fscanf(in, "%d", &n)
	for kase := 1; kase <= n; kase++ {
		fmt.Fscanf(in, "%d", &k)
		f := solve(k)
		fmt.Fprintf(out, "Case #%d: %d = %d * %d = %d * %d\n", kase, k, f[0][0], f[0][1], f[1][0], f[1][1])
	}
}

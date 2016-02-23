// UVa 382 - Perfection

package main

import (
	"fmt"
	"os"
)

func factorSum(n int) string {
	total := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 && i != n { // DEFICIENT if n is 1
			total += i
			if i != 1 && i*i != n {
				total += n / i
			}
		}
	}
	switch {
	case total == n:
		return "PERFECT"
	case total < n:
		return "DEFICIENT"
	default:
		return "ABUNDANT"
	}
}

func main() {
	in, _ := os.Open("382.in")
	defer in.Close()
	out, _ := os.Create("382.out")
	defer out.Close()

	var n int
	fmt.Fprintln(out, "PERFECTION OUTPUT")
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintf(out, "%5d  %s\n", n, factorSum(n))
	}
	fmt.Fprintln(out, "END OF OUTPUT")
}

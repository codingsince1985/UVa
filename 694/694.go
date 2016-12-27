// UVa 694 - The Collatz Sequence

package main

import (
	"fmt"
	"os"
)

func solve(a, l int) int {
	count := 0
	for {
		if a > l {
			break
		}
		count++
		if a == 1 {
			break
		}
		if a%2 == 0 {
			a /= 2
		} else {
			a = 3*a + 1
		}
	}
	return count
}

func main() {
	in, _ := os.Open("694.in")
	defer in.Close()
	out, _ := os.Create("694.out")
	defer out.Close()

	var a, l int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d", &a, &l); a < 0 {
			break
		}
		fmt.Fprintf(out, "Case %d: A = %d, limit = %d, number of terms = %d\n", kase, a, l, solve(a, l))
	}
}

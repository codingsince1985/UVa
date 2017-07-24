// UVa 10930 - A-Sequence

package main

import (
	"fmt"
	"os"
)

func solve(d int, sequence []int) bool {
	max := sequence[d-1]
	dp := make([]bool, max+1)
	for _, a := range sequence {
		if dp[a] {
			return false
		}
		for i, visited := range dp {
			if i+a <= max && visited {
				dp[i+a] = true
			}
		}
		dp[a] = true
	}
	return true
}

func main() {
	in, _ := os.Open("10930.in")
	defer in.Close()
	out, _ := os.Create("10930.out")
	defer out.Close()

	var d int
	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%d", &d); err != nil {
			break
		}
		fmt.Fprintf(out, "Case #%d:", kase)
		sequence := make([]int, d)
		for i := range sequence {
			fmt.Fscanf(in, "%d", &sequence[i])
			fmt.Fprintf(out, " %d", sequence[i])
		}
		fmt.Fprintln(out)
		if solve(d, sequence) {
			fmt.Fprintln(out, "This is an A-sequence.")
		} else {
			fmt.Fprintln(out, "This is not an A-sequence.")
		}
	}
}

// UVa 711 - Dividing up

package main

import (
	"fmt"
	"os"
)

const max = 20000 * 6

func solve(sum int, values [6]int) bool {
	if sum%2 != 0 {
		return false
	}
	dp := make([]int, max+1)
	for i := 0; i <= values[0]; i++ {
		dp[i] = 1
	}
	for i := 2; i <= 6; i++ {
		for j := sum / 2; j >= 0; j-- {
			if dp[j] != 0 {
				for k := 0; k <= values[i-1]; k++ {
					x := k*i + j
					if x > sum/2 || dp[x] >= i {
						break
					}
					dp[x] = i
				}
			}
		}
	}
	return dp[sum/2] != 0
}

func main() {
	in, _ := os.Open("711.in")
	defer in.Close()
	out, _ := os.Create("711.out")
	defer out.Close()

	var values [6]int
	for kase := 1; ; kase++ {
		var sum int
		for i := range values {
			fmt.Fscanf(in, "%d", &values[i])
			sum += values[i] * (i + 1)
		}
		if sum == 0 {
			break
		}
		if fmt.Fprintf(out, "Collection #%d:\n", kase); solve(sum, values) {
			fmt.Fprintln(out, "Can be divided.")
		} else {
			fmt.Fprintln(out, "Can't be divided.")
		}
		fmt.Fprintln(out)
	}
}

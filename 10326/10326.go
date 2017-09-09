// UVa 10326 - The Polynomial Equation

package main

import (
	"fmt"
	"os"
)

func output(out *os.File, coefficients []int64) {
	var sign byte
	for i := len(coefficients) - 1; i >= 0; i-- {
		if coefficients[i] >= 0 {
			sign = '+'
		} else {
			sign = '-'
			coefficients[i] = -coefficients[i]
		}
		if coefficients[i] != 0 {
			if i != len(coefficients)-1 {
				fmt.Fprintf(out, " %c ", sign)
			}
			if coefficients[i] != 1 {
				fmt.Fprintf(out, "%d", coefficients[i])
			}
			if i >= 1 {
				fmt.Fprint(out, "x")
			}
			if i > 1 {
				fmt.Fprintf(out, "^%d", i)
			}
		} else {
			if i == 0 {
				fmt.Fprint(out, " + 0")
			}
		}
	}
	fmt.Fprintln(out, " = 0")
}

func solve(n int, roots []int64) []int64 {
	coefficients := make([]int64, n+1)
	coefficients[0] = 1
	for i, root := range roots {
		dp := make([]int64, n+1)
		for j := 0; j <= i; j++ {
			dp[j] += coefficients[j] * -root
			dp[j+1] += coefficients[j]
		}
		copy(coefficients, dp)
	}
	return coefficients
}

func main() {
	in, _ := os.Open("10326.in")
	defer in.Close()
	out, _ := os.Create("10326.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		roots := make([]int64, n)
		for i := range roots {
			fmt.Fscanf(in, "%d", &roots[i])
		}
		output(out, solve(n, roots))
	}
}

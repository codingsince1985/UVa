// UVa 884 - Factorial Factors

package main

import (
	"fmt"
	"os"
)

const max = 1000001

var factors = func() []int {
	factors := make([]int, max)
	for i := 0; i < max; i++ {
		factors[i] = 1
	}
	factors[1] = 0
	for i := 2; i < max; i++ {
		if factors[i] == 1 {
			for j := 2; i*j < max; j++ {
				factors[i*j] = factors[i] + factors[j]
			}
		}
	}
	for i := 2; i < max; i++ {
		factors[i] += factors[i-1]
	}
	return factors
}()

func main() {
	in, _ := os.Open("884.in")
	defer in.Close()
	out, _ := os.Create("884.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintln(out, factors[n])
	}
}

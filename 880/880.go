// UVa 880 - Cantor Fractions

package main

import (
	"fmt"
	"os"
)

func solve(n int) (int, int) {
	var i, sum int
	for sum < n {
		i++
		sum += i
	}
	return 1 + (sum - n), i - (sum - n)
}

func main() {
	in, _ := os.Open("880.in")
	defer in.Close()
	out, _ := os.Create("880.out")
	defer out.Close()

	var n int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		up, down := solve(n)
		fmt.Fprintf(out, "%d/%d\n", up, down)
	}
}

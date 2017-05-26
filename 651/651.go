// UVa 651 - Deck

package main

import (
	"fmt"
	"os"
)

func solve(n int) float64 {
	var l float64
	for i := 1; i <= n; i++ {
		l += .5 / float64(i)
	}
	return l
}

func main() {
	in, _ := os.Open("651.in")
	defer in.Close()
	out, _ := os.Create("651.out")
	defer out.Close()

	var n int
	fmt.Fprintln(out, "# Cards Overhang")
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fprintf(out, "%5d%10.3f\n", n, solve(n))
	}
}

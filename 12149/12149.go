// UVa 12149 - Feynman

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("12149.in")
	defer in.Close()
	out, _ := os.Create("12149.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintln(out, n*(n+1)*(2*n+1)/6)
	}
}

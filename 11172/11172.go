// UVa 11172 - Relational Operator

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("11172.in")
	defer in.Close()
	out, _ := os.Create("11172.out")
	defer out.Close()

	var n, n1, n2 int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d", &n1, &n2)
		switch {
		case n1 > n2:
			fmt.Fprintln(out, ">")
		case n1 < n2:
			fmt.Fprintln(out, "<")
		default:
			fmt.Fprintln(out, "=")
		}
	}
}

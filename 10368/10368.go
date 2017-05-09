// UVa 10368 - Euclid's Game

package main

import (
	"fmt"
	"os"
)

func dfs(n1, n2 int) bool {
	if n1 < n2 {
		n1, n2 = n2, n1
	}
	switch {
	case n1%n2 == 0:
		return true
	case n1/n2 == 1:
		return !dfs(n2, n1%n2)
	default:
		return !dfs(n1%n2+n2, n2) || !dfs(n2, n1%n2)
	}
}

func main() {
	in, _ := os.Open("10368.in")
	defer in.Close()
	out, _ := os.Create("10368.out")
	defer out.Close()

	var n1, n2 int
	for {
		if fmt.Fscanf(in, "%d%d", &n1, &n2); n1 == 0 && n2 == 0 {
			break
		}
		if dfs(n1, n2) {
			fmt.Fprintln(out, "Stan wins")
		} else {
			fmt.Fprintln(out, "Ollie wins")
		}
	}
}

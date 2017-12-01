// UVa 10120 - Gift?!

package main

import (
	"fmt"
	"os"
)

var n, m int

func dfs(curr, step int) bool {
	switch {
	case curr < 1 || curr > n:
		return false
	case curr == m:
		return true
	default:
		return dfs(curr+2*step-1, step+1) || dfs(curr-2*step+1, step+1)
	}
}

func main() {
	in, _ := os.Open("10120.in")
	defer in.Close()
	out, _ := os.Create("10120.out")
	defer out.Close()

	for {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		if n >= 49 || dfs(1, 2) {
			fmt.Fprintln(out, "Let me try!")
		} else {
			fmt.Fprintln(out, "Don't make fun of me!")
		}
	}
}

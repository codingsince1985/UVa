// UVa 278 - Chess

package main

import (
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	in, _ := os.Open("278.in")
	defer in.Close()
	out, _ := os.Create("278.out")
	defer out.Close()

	var kase, m, n int
	var p byte
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanf(in, "%c%d%d", &p, &m, &n)
		switch p {
		case 'r', 'Q':
			fmt.Fprintln(out, min(m, n))
		case 'k':
			fmt.Fprintln(out, (m*n+1)/2)
		case 'K':
			fmt.Fprintln(out, (m+1)/2*((n+1)/2))
		}
		kase--
	}
}

// UVa 278 - Chess

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("278.in")
	defer in.Close()
	out, _ := os.Create("278.out")
	defer out.Close()

	var kase, m, n int
	var p byte
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		switch fmt.Fscanf(in, "%c%d%d", &p, &m, &n); p {
		case 'r', 'Q':
			fmt.Fprintln(out, min(m, n))
		case 'k':
			fmt.Fprintln(out, (m*n+1)/2)
		case 'K':
			fmt.Fprintln(out, (m+1)/2*((n+1)/2))
		}
	}
}

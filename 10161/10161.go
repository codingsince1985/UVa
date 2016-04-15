// UVa 10161 - Ant on a Chessboard

package main

import (
	"fmt"
	"os"
)

func findSquare(n int) int {
	for i := 1; i <= 44722; i++ {
		if i*i >= n {
			return i
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("10161.in")
	defer in.Close()
	out, _ := os.Create("10161.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		s := findSquare(n)
		var col, row int
		if s%2 == 0 {
			if n >= (s-1)*(s-1)+s {
				col, row = s, s*s-n+1
			} else {
				col, row = n-(s-1)*(s-1), s
			}
		} else {
			if n >= (s-1)*(s-1)+s {
				col, row = s*s-n+1, s
			} else {
				col, row = s, n-(s-1)*(s-1)
			}
		}
		fmt.Fprintln(out, col, row)
	}
}

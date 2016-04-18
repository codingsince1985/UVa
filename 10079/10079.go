// UVa 10079 - Pizza Cutting

package main

import (
	"fmt"
	"os"
)

func cut(n int64) int64 { return n*(n+1)/2 + 1 }

func main() {
	in, _ := os.Open("10079.in")
	defer in.Close()
	out, _ := os.Create("10079.out")
	defer out.Close()

	var n int64
	for {
		if fmt.Fscanf(in, "%d", &n); n < 0 {
			break
		}
		fmt.Fprintln(out, cut(n))
	}
}

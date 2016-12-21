// UVa 10489 - Boxes of Chocolates

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10489.in")
	defer in.Close()
	out, _ := os.Create("10489.out")
	defer out.Close()

	var t, n, b, k, num int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		total := 0
		for fmt.Fscanf(in, "%d%d", &n, &b); b > 0; b-- {
			// (a1 * a2 * ... * ak) % m = (((a1 % m) * a2) % m) * a3) ... ) % m
			remaining := 1
			for fmt.Fscanf(in, "%d", &k); k > 0; k-- {
				fmt.Fscanf(in, "%d", &num)
				remaining = (remaining * num) % n
			}
			total = (total + remaining) % n
		}
		fmt.Fprintln(out, total)
	}
}

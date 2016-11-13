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
	fmt.Fscanf(in, "%d", &t)
	for t > 0 {
		fmt.Fscanf(in, "%d%d", &n, &b)
		total := 0
		for b > 0 {
			fmt.Fscanf(in, "%d", &k)
			// (a1 * a2 * ... * ak) % m = (((a1 % m) * a2) % m) * a3) ... ) % m
			remaining := 1
			for k > 0 {
				fmt.Fscanf(in, "%d", &num)
				remaining = (remaining * num) % n
				k--
			}
			total = (total + remaining) % n
			b--
		}
		fmt.Fprintln(out, total)
		t--
	}
}

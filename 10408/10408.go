// UVa 10408 - Farey sequences

package main

import (
	"fmt"
	"os"
)

func farey(n, k int) (int, int) {
	a0, b0, a1, b1 := 0, 1, 1, n
	for k > 1 {
		r := (n + b0) / b1
		a0, b0, a1, b1 = a1, b1, r*a1-a0, r*b1-b0
		k--
	}
	return a1, b1
}

func main() {
	in, _ := os.Open("10408.in")
	defer in.Close()
	out, _ := os.Create("10408.out")
	defer out.Close()

	var n, k int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &k); err != nil {
			break
		}
		a, b := farey(n, k)
		fmt.Fprintf(out, "%d/%d\n", a, b)
	}
}

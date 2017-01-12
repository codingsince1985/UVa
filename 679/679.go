// UVa 679 - Dropping Balls

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("679.in")
	defer in.Close()
	out, _ := os.Create("679.out")
	defer out.Close()

	var n, D, I int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d", &D, &I)
		k := 1
		for i := 0; i < D-1; i++ {
			if I%2 == 0 {
				k = 2*k + 1
				I = I / 2
			} else {
				k = 2 * k
				I = (I + 1) / 2
			}
		}
		fmt.Fprintln(out, k)
	}
}

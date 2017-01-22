// UVa 10673 - Play with Floor and Ceil

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10673.in")
	defer in.Close()
	out, _ := os.Create("10673.out")
	defer out.Close()

	var t, x, k int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d%d", &x, &k)
		if x%k != 0 {
			fmt.Fprintln(out, k-x%k, x%k)
		} else {
			fmt.Fprintln(out, 0, k)
		}
	}
}

// UVa 10970 - Big Chocolate

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10970.in")
	defer in.Close()
	out, _ := os.Create("10970.out")
	defer out.Close()

	var m, n int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &m, &n); err != nil {
			break
		}
		fmt.Fprintln(out, m*n-1)
	}
}

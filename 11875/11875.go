// UVa 11875 - Brick Game

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("11875.in")
	defer in.Close()
	out, _ := os.Create("11875.out")
	defer out.Close()

	var kase, n int
	fmt.Fscanf(in, "%d", &kase)
	for i := 1; i <= kase; i++ {
		fmt.Fscanf(in, "%d", &n)
		ages := make([]int, n)
		for i := range ages {
			fmt.Fscanf(in, "%d", &ages[i])
		}
		fmt.Fprintf(out, "Case %d: %d\n", i, ages[n/2])
	}
}

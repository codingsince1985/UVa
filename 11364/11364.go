// UVa 11364 - Parking

package main

import (
	"fmt"
	"os"
	"sort"
)

func solve(n int, shops []int) int {
	sort.Ints(shops)
	return (shops[n-1] - shops[0]) * 2
}

func main() {
	in, _ := os.Open("11364.in")
	defer in.Close()
	out, _ := os.Create("11364.out")
	defer out.Close()

	var kase, n int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d", &n)
		shops := make([]int, n)
		for i := range shops {
			fmt.Fscanf(in, "%d", &shops[i])
		}
		fmt.Fprintln(out, solve(n, shops))
	}
}

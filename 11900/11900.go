// UVa 11900 - Boiled Eggs

package main

import (
	"fmt"
	"os"
)

func solve(eggs []int, p, q int) int {
	var num, weight int
	for _, egg := range eggs {
		if weight += egg; weight > q || num == p {
			break
		}
		num++
	}
	return num
}

func main() {
	in, _ := os.Open("11900.in")
	defer in.Close()
	out, _ := os.Create("11900.out")
	defer out.Close()

	var t, n, p, q int
	fmt.Fscanf(in, "%d", &t)
	for kase := 1; kase <= t; kase++ {
		fmt.Fscanf(in, "%d%d%d", &n, &p, &q)
		eggs := make([]int, n)
		for i := range eggs {
			fmt.Fscanf(in, "%d", &eggs[i])
		}
		fmt.Fprintf(out, "Case %d: %d\n", kase, solve(eggs, p, q))
	}
}

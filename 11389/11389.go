// UVa 11389 - The Bus Driver Problem

package main

import (
	"fmt"
	"os"
	"sort"
)

func solve(d, r int, morning, evening []int) int {
	sort.Ints(morning)
	sort.Sort(sort.Reverse(sort.IntSlice(evening)))
	var amount int
	for i := range morning {
		amount += (morning[i] + evening[i] - d) * r
	}
	return amount
}

func main() {
	in, _ := os.Open("11389.in")
	defer in.Close()
	out, _ := os.Create("11389.out")
	defer out.Close()

	var n, d, r int
	for {
		if fmt.Fscanf(in, "%d%d%d", &n, &d, &r); n == 0 && d == 0 && r == 0 {
			break
		}
		morning := make([]int, n)
		for i := range morning {
			fmt.Fscanf(in, "%d", &morning[i])
		}
		evening := make([]int, n)
		for i := range evening {
			fmt.Fscanf(in, "%d", &evening[i])
		}
		fmt.Fprintln(out, solve(d, r, morning, evening))
	}
}

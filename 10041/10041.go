// UVa 10041 - Vito's Family

package main

import (
	"fmt"
	"os"
	"sort"
)

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func solve(lst []int) int {
	sort.Ints(lst)
	total := 0
	median := lst[len(lst)/2]
	for _, v := range lst {
		total += abs(v - median)
	}
	return total
}

func main() {
	in, _ := os.Open("10041.in")
	defer in.Close()
	out, _ := os.Create("10041.out")
	defer out.Close()

	var n, r, tmp int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		var address []int
		for fmt.Fscanf(in, "%d", &r); r > 0; r-- {
			fmt.Fscanf(in, "%d", &tmp)
			address = append(address, tmp)
		}
		fmt.Fprintln(out, solve(address))
	}
}

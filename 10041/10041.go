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
	fmt.Fscanf(in, "%d", &n)
	for n > 0 {
		var address []int
		fmt.Fscanf(in, "%d", &r)
		for r > 0 {
			fmt.Fscanf(in, "%d", &tmp)
			address = append(address, tmp)
			r--
		}
		sort.Ints(address)
		fmt.Fprintln(out, solve(address))
		n--
	}
}

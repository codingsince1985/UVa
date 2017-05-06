// UVa 11369 - Shopaholic

package main

import (
	"fmt"
	"os"
	"sort"
)

func solve(prices []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(prices)))
	var total int
	for i, price := range prices {
		if i%3 == 2 {
			total += price
		}
	}
	return total
}

func main() {
	in, _ := os.Open("11369.in")
	defer in.Close()
	out, _ := os.Create("11369.out")
	defer out.Close()

	var t, n int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d", &n)
		prices := make([]int, n)
		for i := range prices {
			fmt.Fscanf(in, "%d", &prices[i])
		}
		fmt.Fprintln(out, solve(prices))
	}
}

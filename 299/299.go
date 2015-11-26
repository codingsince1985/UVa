// UVa 299 - Train Swapping

package main

import (
	"fmt"
	"os"
)

func swap(t []int) int {
	c := 0
	l := len(t)
	for i, v := range t[:l - 1] {
		for _, w := range t[i + 1:] {
			if v > w {
				c++
			}
		}
	}
	return c
}

func main() {
	in, _ := os.Open("299.in")
	out, _ := os.Create("299.out")
	var n, m int
	var t []int
	fmt.Fscanf(in, "%d", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(in, "%d", &m)
		t = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscanf(in, "%d", &t[j])
		}
		fmt.Fprintf(out, "Optimal train swapping takes %d swaps.\n", swap(t))
	}
	in.Close()
	out.Close()
}

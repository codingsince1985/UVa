// UVa 299 - Train Swapping

package main

import (
	"fmt"
	"os"
)

func swap(t []int) int {
	c := 0
	for i, v := range t[:len(t)-1] {
		for _, w := range t[i+1:] {
			if v > w {
				c++
			}
		}
	}
	return c
}

func main() {
	in, _ := os.Open("299.in")
	defer in.Close()
	out, _ := os.Create("299.out")
	defer out.Close()

	var n, m int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d", &m)
		t := make([]int, m)
		for j := range t {
			fmt.Fscanf(in, "%d", &t[j])
		}
		fmt.Fprintf(out, "Optimal train swapping takes %d swaps.\n", swap(t))
	}
}

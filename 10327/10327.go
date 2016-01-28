// UVa 10327 - Flip Sort

package main

import (
	"fmt"
	"os"
)

func swap(t []int) int {
	c := 0
	l := len(t)
	for i, v := range t[:l-1] {
		for _, w := range t[i+1:] {
			if v > w {
				c++
			}
		}
	}
	return c
}

func main() {
	in, _ := os.Open("10327.in")
	defer in.Close()
	out, _ := os.Create("10327.out")
	defer out.Close()

	var n int
	var a []int
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		a = make([]int, n)
		for i := range a {
			fmt.Fscanf(in, "%d", &a[i])
		}
		fmt.Fprintf(out, "Minimum exchange operations : %d\n", swap(a))
	}
}

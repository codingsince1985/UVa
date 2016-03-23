// UVa 305 - Joseph

package main

import (
	"fmt"
	"os"
)

func josephus(k, m int) int {
	var i int
	size := 2 * k
	current := (m - 1) % size
	for i = 0; i < 2*k; i++ {
		if current < k {
			break
		}
		size--
		current = (current + m - 1) % size
	}
	return i
}

func main() {
	in, _ := os.Open("305.in")
	defer in.Close()
	out, _ := os.Create("305.out")
	defer out.Close()

	var k int
	for {
		if fmt.Fscanf(in, "%d", &k); k == 0 {
			break
		}
		for m := k + 1; ; m++ {
			if josephus(k, m) == k {
				fmt.Fprintln(out, m)
				break
			}
		}
	}
}

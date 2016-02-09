// UVa 136 - Ugly Numbers

package main

import (
	"fmt"
	"os"
)

const MAX = 1500

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	out, _ := os.Create("136.out")
	defer out.Close()

	var lst [MAX + 1]int
	lst[1] = 1
	current, p2, p3, p5 := 1, 1, 1, 1
	for current < MAX {
		for lst[p2]*2 <= lst[current] {
			p2++
		}
		for lst[p3]*3 <= lst[current] {
			p3++
		}
		for lst[p5]*5 <= lst[current] {
			p5++
		}
		current++
		lst[current] = min(lst[p2]*2, min(lst[p3]*3, lst[p5]*5))
	}
	fmt.Fprintf(out, "The 1500'th ugly number is %d.\n", lst[1500])
}

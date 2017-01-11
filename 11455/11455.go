// UVa 11455 - Behold my quadrangle

package main

import (
	"fmt"
	"os"
	"sort"
)

func solve(sides []int) string {
	sort.Ints(sides)
	if sides[0] == sides[1] && sides[2] == sides[3] {
		if sides[1] == sides[2] {
			return "square"
		}
		return "rectangle"
	}
	if sides[3] > sides[0]+sides[1]+sides[2] {
		return "banana"
	}
	return "quadrangle"
}

func main() {
	in, _ := os.Open("11455.in")
	defer in.Close()
	out, _ := os.Create("11455.out")
	defer out.Close()

	var kase int
	sides := make([]int, 4)
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d%d%d", &sides[0], &sides[1], &sides[2], &sides[3])
		fmt.Fprintln(out, solve(sides))
	}
}

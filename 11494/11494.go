// UVa 11494 - Queen

package main

import (
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func solve(x1, y1, x2, y2 int) int {
	switch {
	case x1 == x2 && y1 == y2:
		return 0
	case x1 == x2 || y1 == y2 || abs(x1-x2) == abs(y1-y2):
		return 1
	default:
		return 2
	}
}

func main() {
	in, _ := os.Open("11494.in")
	defer in.Close()
	out, _ := os.Create("11494.out")
	defer out.Close()

	var x1, x2, y1, y2 int
	for {
		if fmt.Fscanf(in, "%d%d%d%d", &x1, &y1, &x2, &y2); x1 == 0 && y1 == 0 && x2 == 0 && y2 == 0 {
			break
		}
		fmt.Fprintln(out, solve(x1, y1, x2, y2))
	}
}

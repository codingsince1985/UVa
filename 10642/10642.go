// UVa 10642 - Can You Solve It?

package main

import (
	"fmt"
	"os"
)

func distance(x, y int) int { return (1+x+y)*(x+y)/2 + x }

func walk(x1, y1, x2, y2 int) int { return distance(x2, y2) - distance(x1, y1) }

func main() {
	in, _ := os.Open("10642.in")
	defer in.Close()
	out, _ := os.Create("10642.out")
	defer out.Close()

	var kase, x1, y1, x2, y2 int
	fmt.Fscanf(in, "%d", &kase)
	for i := 1; i <= kase; i++ {
		fmt.Fscanf(in, "%d%d%d%d", &x1, &y1, &x2, &y2)
		fmt.Fprintf(out, "Case %d: %d\n", i, walk(x1, y1, x2, y2))
	}
}

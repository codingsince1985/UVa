// UVa 10865 - Brownie Points

package main

import (
	"fmt"
	"os"
)

type brownie struct{ x, y int }

func solve(n int, brownies []brownie) (int, int) {
	var stan, ollie int
	center := brownies[(n-1)/2]
	for _, b := range brownies {
		switch pos := (b.x - center.x) * (b.y - center.y); {
		case pos > 0:
			stan++
		case pos < 0:
			ollie++
		}
	}
	return stan, ollie
}

func main() {
	in, _ := os.Open("10865.in")
	defer in.Close()
	out, _ := os.Create("10865.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		brownies := make([]brownie, n)
		for i := range brownies {
			fmt.Fscanf(in, "%d%d", &brownies[i].x, &brownies[i].y)
		}
		stan, ollie := solve(n, brownies)
		fmt.Fprintln(out, stan, ollie)
	}
}

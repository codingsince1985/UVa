// UVa 10167 - Birthday Cake

package main

import (
	"fmt"
	"os"
)

type cherry struct{ x, y int }

func solve(n int, cherries []cherry) (bool, int, int) {
	for a := -500; a <= 500; a++ {
		for b := -500; b <= 500; b++ {
			if a != 0 || b != 0 {
				var lucy, lily int
				for _, c := range cherries {
					r := a*c.x + b*c.y
					switch {
					case r > 0:
						lucy++
					case r < 0:
						lily++
					}
				}
				if lucy == lily && lucy == n {
					return true, a, b
				}
			}
		}
	}
	return false, 0, 0
}

func main() {
	in, _ := os.Open("10167.in")
	defer in.Close()
	out, _ := os.Create("10167.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		cherries := make([]cherry, 2*n)
		for i := range cherries {
			fmt.Fscanf(in, "%d%d", &cherries[i].x, &cherries[i].y)
		}
		if ok, a, b := solve(n, cherries); ok {
			fmt.Fprintln(out, a, b)
		}
	}
}

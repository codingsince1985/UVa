// UVa 10365 - Blocks

package main

import (
	"fmt"
	"math"
	"os"
)

func solve(n int) int {
	min := math.MaxInt32
	for x := 1; x <= int(math.Sqrt(float64(n))); x++ {
		if n%x != 0 {
			continue
		}
		y := n / x
		for z := 1; z <= int(math.Sqrt(float64(y))); z++ {
			if y%z != 0 {
				continue
			}
			if tmp := 2 * (x*z + y + x*y/z); tmp < min {
				min = tmp
			}
		}
	}
	return min
}

func main() {
	in, _ := os.Open("10365.in")
	defer in.Close()
	out, _ := os.Create("10365.out")
	defer out.Close()

	var c, n int
	for fmt.Fscanf(in, "%d", &c); c > 0; c-- {
		fmt.Fscanf(in, "%d", &n)
		fmt.Fprintln(out, solve(n))
	}
}

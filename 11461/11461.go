// UVa 11461 - Square Numbers

package main

import (
	"fmt"
	"math"
	"os"
)

func solve(a, b int) int {
	sqrta, sqrtb := math.Sqrt(float64(a)), math.Sqrt(float64(b))
	cnt := int(sqrtb) - int(sqrta)
	if sqrta*sqrta == float64(a) {
		cnt++
	}
	return cnt
}

func main() {
	in, _ := os.Open("11461.in")
	defer in.Close()
	out, _ := os.Create("11461.out")
	defer out.Close()

	var a, b int
	for {
		if fmt.Fscanf(in, "%d%d", &a, &b); a == 0 && b == 0 {
			break
		}
		fmt.Fprintln(out, solve(a, b))
	}
}

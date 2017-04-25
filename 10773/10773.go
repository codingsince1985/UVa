// UVa 10773 - Back to Intermediate Math

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("10773.in")
	defer in.Close()
	out, _ := os.Create("10773.out")
	defer out.Close()

	var kase int
	fmt.Fscanf(in, "%d", &kase)
	var d, v, u float64
	for i := 1; i <= kase; i++ {
		fmt.Fprintf(out, "Case %d: ", i)
		if fmt.Fscanf(in, "%f%f%f", &d, &v, &u); v >= u {
			fmt.Fprintln(out, "can’t determine")
		} else {
			fmt.Fprintf(out, "%.3f\n", d/math.Sqrt(u*u-v*v)-d/u)
		}
	}
}

// UVa 10209 - Is This Integration ?

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("10209.in")
	defer in.Close()
	out, _ := os.Create("10209.out")
	defer out.Close()

	sqrt3 := math.Sqrt(3)
	xr := 1 - sqrt3 + math.Pi/3
	yr := 2*sqrt3 - 4 + math.Pi/3
	zr := 1 - xr - yr

	var a float64
	for {
		if _, err := fmt.Fscanf(in, "%f", &a); err != nil {
			break
		}
		area := a * a
		x := area * xr
		y := area * yr
		z := area * zr
		fmt.Fprintf(out, "%.3f %.3f %.3f\n", x, y, z)
	}
}

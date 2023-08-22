// UVa 10215 - The Largest/Smallest Box ...

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("10215.in")
	defer in.Close()
	out, _ := os.Create("10215.out")
	defer out.Close()

	var l, w float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f", &l, &w); err != nil {
			break
		}
		fmt.Fprintf(out, "%.3f %.3f %.3f\n", (l+w-math.Sqrt(l*l-l*w+w*w))/6, 0.0, min(l, w)/2)
	}
}

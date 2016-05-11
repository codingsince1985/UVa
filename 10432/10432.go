// UVa 10432 - Polygon Inside A Circle

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("10432.in")
	defer in.Close()
	out, _ := os.Create("10432.out")
	defer out.Close()

	var r, n float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f", &r, &n); err != nil {
			break
		}
		fmt.Fprintf(out, "%.3f\n", n*math.Sin(2*math.Pi/n)*r*r/2)
	}
}

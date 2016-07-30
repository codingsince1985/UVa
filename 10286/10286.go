// UVa 10286 - Trouble with a Pentagon

package main

import (
	"fmt"
	"math"
	"os"
)

func sin(degree float64) float64 { return math.Sin(degree * math.Pi / 180) }

func main() {
	in, _ := os.Open("10286.in")
	defer in.Close()
	out, _ := os.Create("10286.out")
	defer out.Close()

	var f float64
	for {
		if _, err := fmt.Fscanf(in, "%f", &f); err != nil {
			break
		}
		fmt.Fprintf(out, "%.10f\n", f*sin(108)/sin(63))
	}
}

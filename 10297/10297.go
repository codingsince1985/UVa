// UVa 10297 - Beavergnaw

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("10297.in")
	defer in.Close()
	out, _ := os.Create("10297.out")
	defer out.Close()

	var d, v float64
	for {
		if fmt.Fscanf(in, "%f%f", &d, &v); d == 0 && v == 0 {
			break
		}
		fmt.Fprintf(out, "%.3f\n", math.Pow(d*d*d-(6*v)/math.Pi, 1.0/3))
	}
}

// UVa 438 - The Circumference of the Circle

package main

import (
	"fmt"
	"math"
	"os"
)

type point struct{ x, y float64 }

func distance(p1, p2 point) float64 { return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)) }

func radius(d1, d2, d3 float64) float64 {
	t := d1 + d2 + d3
	return d1 * d2 * d3 / math.Sqrt(t*(t-2*d1)*(t-2*d2)*(t-2*d3))
}

func main() {
	in, _ := os.Open("438.in")
	defer in.Close()
	out, _ := os.Create("438.out")
	defer out.Close()

	var p1, p2, p3 point
	for {
		if _, err := fmt.Fscanf(in, "%f%f%f%f%f%f", &p1.x, &p1.y, &p2.x, &p2.y, &p3.x, &p3.y); err != nil {
			break
		}
		d1, d2, d3 := distance(p1, p2), distance(p2, p3), distance(p3, p1)
		fmt.Fprintf(out, "%.2f\n", radius(d1, d2, d3)*2*math.Pi)
	}
}

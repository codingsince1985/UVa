// UVa 10310 - Dog and Gopher

package main

import (
	"fmt"
	"math"
	"os"
)

type point struct {
	x, y float64
}

func distance(p1, p2 point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func main() {
	in, _ := os.Open("10310.in")
	defer in.Close()
	out, _ := os.Create("10310.out")
	defer out.Close()

	var n int
	var x, y float64
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fscanf(in, "%f%f", &x, &y)
		gopher := point{x, y}
		fmt.Fscanf(in, "%f%f", &x, &y)
		dog := point{x, y}

		found := false
		var p point
		for n > 0 {
			fmt.Fscanf(in, "%f%f", &x, &y)
			p = point{x, y}
			if distance(p, gopher)*2 <= distance(p, dog) {
				found = true
				break
			}
			n--
		}
		if !found {
			fmt.Fprintln(out, "The gopher cannot escape.")
		} else {
			fmt.Fprintf(out, "The gopher can escape through the hole at (%.3f,%.3f).\n", p.x, p.y)
		}
		fmt.Fscanln(in)
	}
}

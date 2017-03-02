// UVa 10310 - Dog and Gopher

package main

import (
	"fmt"
	"math"
	"os"
)

type point struct{ x, y float64 }

func distance(p1, p2 point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func main() {
	in, _ := os.Open("10310.in")
	defer in.Close()
	out, _ := os.Create("10310.out")
	defer out.Close()

	var n int
	var gopher, dog, p point
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		fmt.Fscanf(in, "%f%f", &gopher.x, &gopher.y)
		fmt.Fscanf(in, "%f%f", &dog.x, &dog.y)

		found := false
		for ; n > 0; n-- {
			fmt.Fscanf(in, "%f%f", &p.x, &p.y)
			if distance(p, gopher)*2 <= distance(p, dog) {
				found = true
				break
			}
		}
		if !found {
			fmt.Fprintln(out, "The gopher cannot escape.")
		} else {
			fmt.Fprintf(out, "The gopher can escape through the hole at (%.3f,%.3f).\n", p.x, p.y)
		}
		fmt.Fscanln(in)
	}
}

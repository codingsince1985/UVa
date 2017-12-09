// UVa 356 - Square Pegs And Round Holes

package main

import (
	"fmt"
	"math"
	"os"
)

func distance(x, y int) float64 { return math.Sqrt(float64(x*x + y*y)) }

func count(n int) (int, int) {
	var whollyIn, halfIn int
	radius := float64(n) - 0.5
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			if distance(x, y) < radius {
				if distance(x+1, y+1) <= radius {
					whollyIn++
				} else {
					halfIn++
				}
			}
		}
	}
	return whollyIn, halfIn
}

func main() {
	in, _ := os.Open("356.in")
	defer in.Close()
	out, _ := os.Create("356.out")
	defer out.Close()

	var n int
	for first := true; ; {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		whollyIn, halfIn := count(n)
		fmt.Fprintf(out, "In the case n = %d, %d cells contain segments of the circle.\n", n, halfIn*4)
		fmt.Fprintf(out, "There are %d cells completely contained in the circle.\n", whollyIn*4)
	}
}

// UVa 10233 - Dermuba Triangle

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	height        = math.Cos(math.Pi / 6)
	partialHeight = math.Tan(math.Pi/6) / 2
)

type point struct{ x, y float64 }

func distance(p1, p2 point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func pinPoint(n int) point {
	level := int(math.Sqrt(float64(n)))
	offset := n - level*level + 1
	x := float64(offset-level-1) / 2
	y := float64(level) * height
	if offset%2 == 1 {
		y += height - partialHeight
	} else {
		y += partialHeight
	}
	return point{x, y}
}

func main() {
	in, _ := os.Open("10233.in")
	defer in.Close()
	out, _ := os.Create("10233.out")
	defer out.Close()

	var n, m int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &n, &m); err != nil {
			break
		}
		fmt.Fprintf(out, "%.3f\n", distance(pinPoint(n), pinPoint(m)))
	}
}

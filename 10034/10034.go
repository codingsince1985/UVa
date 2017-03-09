// UVa 10034 - Freckles

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

func prim(freckles []point) float64 {
	visited := make(map[int]bool)
	visited[0] = true
	var total float64
	for i := range freckles {
		if !visited[i] {
			var minP int
			minD := math.MaxFloat64
			for j := range visited {
				if dis := distance(freckles[i], freckles[j]); dis < minD {
					minP = i
					minD = dis
				}
			}
			visited[minP] = true
			total += minD
		}
	}
	return total
}

func main() {
	in, _ := os.Open("10034.in")
	defer in.Close()
	out, _ := os.Create("10034.out")
	defer out.Close()

	var kase, num int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &num)
		freckles := make([]point, num)
		for i := range freckles {
			fmt.Fscanf(in, "%f%f", &freckles[i].x, &freckles[i].y)
		}
		fmt.Fprintf(out, "%.2f\n", prim(freckles))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}

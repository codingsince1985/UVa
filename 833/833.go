// UVa 833 - Water Falls

package main

import (
	"fmt"
	"io"
	"os"
)

type (
	point struct{ x, y int }
	line  struct{ p1, p2 point }
)

func inBetween(p point, l line) bool { return (p.x-l.p1.x)*(p.x-l.p2.x) <= 0 }

func touchAt(p point, l line) float64 {
	return float64(l.p1.y) + (float64(p.x-l.p1.x) / float64(l.p2.x-l.p1.x) * float64(l.p2.y-l.p1.y))
}

func solve(out io.Writer, lines []line, points []point) {
	for _, p := range points {
		for {
			highest := -1.0
			var idx int
			for i, l := range lines {
				if inBetween(p, l) {
					if touch := touchAt(p, l); touch < float64(p.y) && touch > highest {
						highest, idx = touch, i
					}
				}
			}
			if highest == -1 {
				fmt.Fprintln(out, p.x)
				break
			}
			if lines[idx].p2.y > lines[idx].p1.y {
				p = lines[idx].p1
			} else {
				p = lines[idx].p2
			}
		}
	}
}

func main() {
	in, _ := os.Open("833.in")
	defer in.Close()
	out, _ := os.Create("833.out")
	defer out.Close()

	var kase, np, ns int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &np)
		lines := make([]line, np)
		for i := range lines {
			fmt.Fscanf(in, "%d%d%d%d", &lines[i].p1.x, &lines[i].p1.y, &lines[i].p2.x, &lines[i].p2.y)
		}
		fmt.Fscanf(in, "%d", &ns)
		points := make([]point, ns)
		for i := range points {
			fmt.Fscanf(in, "%d%d", &points[i].x, &points[i].y)
		}
		solve(out, lines, points)
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}

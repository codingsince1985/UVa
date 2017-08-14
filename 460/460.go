// UVa 460 - Overlapping Rectangles

package main

import (
	"fmt"
	"os"
)

type (
	point  struct{ x, y int }
	window struct{ p1, p2 point }
)

func in(p point, w window) bool { return p.x > w.p1.x && p.x < w.p2.x && p.y > w.p1.y && p.y < w.p2.y }

func solve(w1, w2 window) *window {
	switch {
	case in(w1.p1, w2):
		return &window{w1.p1, w2.p2}
	case in(w1.p2, w2):
		return &window{w2.p1, w1.p2}
	case in(point{w1.p1.x, w1.p2.y}, w2):
		return &window{point{w1.p1.x, w2.p1.y}, point{w2.p2.x, w1.p2.y}}
	case in(point{w1.p2.x, w1.p1.y}, w2):
		return &window{point{w2.p1.x, w1.p1.y}, point{w1.p2.x, w2.p2.y}}
	default:
		return nil
	}
}

func main() {
	in, _ := os.Open("460.in")
	defer in.Close()
	out, _ := os.Create("460.out")
	defer out.Close()

	var kase int
	var w1, w2 window
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d%d%d%d", &w1.p1.x, &w1.p1.y, &w1.p2.x, &w1.p2.y)
		fmt.Fscanf(in, "%d%d%d%d", &w2.p1.x, &w2.p1.y, &w2.p2.x, &w2.p2.y)
		if w := solve(w1, w2); w == nil {
			fmt.Fprintln(out, "No Overlap")
		} else {
			fmt.Fprintf(out, "%d %d %d %d\n", w.p1.x, w.p1.y, w.p2.x, w.p2.y)
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}

// UVa 638 - Finding Rectangles

package main

import (
	"fmt"
	"os"
)

type point struct {
	label string
	x, y  int
}

func solve(points []point) []string {
	xMap := make(map[int][]point)
	for _, p := range points {
		xMap[p.x] = append(xMap[p.x], p)
	}
	var rectangles []string
	for _, pi := range points {
		for _, pj := range points {
			if pi.x < pj.x && pi.y == pj.y {
				for _, pk := range xMap[pi.x] {
					if pk.y < pi.y {
						for _, pl := range xMap[pj.x] {
							if pl.y < pj.y && pk.x < pl.x && pk.y == pl.y {
								rectangles = append(rectangles, pi.label+pj.label+pl.label+pk.label)
							}
						}
					}
				}
			}
		}
	}
	return rectangles
}

func main() {
	in, _ := os.Open("638.in")
	defer in.Close()
	out, _ := os.Create("638.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		points := make([]point, n)
		for i := range points {
			fmt.Fscanf(in, "%s%d%d", &points[i].label, &points[i].x, &points[i].y)
		}
		if rectangles := solve(points); len(rectangles) == 0 {
			fmt.Fprintf(out, "Point set %d: No rectangles\n", kase)
		} else {
			fmt.Fprintf(out, "Point set %d:\n", kase)
			for i, rectangle := range rectangles {
				fmt.Fprintf(out, " %s", rectangle)
				if (i+1)%10 == 0 {
					fmt.Fprintln(out)
				}
			}
			if len(rectangles)%10 != 0 {
				fmt.Fprintln(out)
			}
		}
	}
}

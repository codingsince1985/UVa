// UVa 476 - Points in Figures: Rectangles

package main

import (
	"fmt"
	"os"
)

type rectangle struct {
	x1, y1, x2, y2 float64
}

func testIn(out *os.File, count int, x, y float64, rectangles []rectangle) {
	var inAny bool
	for i, v := range rectangles {
		if x > v.x1 && x < v.x2 && y < v.y1 && y > v.y2 {
			fmt.Fprintf(out, "Point %d is contained in figure %d\n", count, i+1)
			inAny = true
		}
	}
	if !inAny {
		fmt.Fprintf(out, "Point %d is not contained in any figure\n", count)
	}
}

func main() {
	in, _ := os.Open("476.in")
	defer in.Close()
	out, _ := os.Create("476.out")
	defer out.Close()

	var rectangles []rectangle
	var a, b, c, d float64
	var str string
	for {
		if fmt.Fscanf(in, "%s", &str); str == "*" {
			break
		}
		fmt.Fscanf(in, "%f%f%f%f", &a, &b, &c, &d)
		rectangles = append(rectangles, rectangle{a, b, c, d})
	}

	count := 0
	for {
		if fmt.Fscanf(in, "%f%f", &a, &b); a == 9999.9 && b == 9999.9 {
			break
		}
		count++
		testIn(out, count, a, b, rectangles)
	}
}

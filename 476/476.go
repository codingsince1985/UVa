// UVa 476 - Points in Figures: Rectangles

package main

import (
	"fmt"
	"os"
)

var out *os.File

type point struct{ x, y float64 }

type rectangle struct{ p1, p2 point }

func testIn(count int, p point, rectangles []rectangle) {
	inAny := false
	for i, v := range rectangles {
		if p.x > v.p1.x && p.x < v.p2.x && p.y < v.p1.y && p.y > v.p2.y {
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
	out, _ = os.Create("476.out")
	defer out.Close()

	var rectangles []rectangle
	var a, b, c, d float64
	var str string
	for {
		if fmt.Fscanf(in, "%s", &str); str == "*" {
			break
		}
		fmt.Fscanf(in, "%f%f%f%f", &a, &b, &c, &d)
		rectangles = append(rectangles, rectangle{point{a, b}, point{c, d}})
	}

	count := 0
	for {
		if fmt.Fscanf(in, "%f%f", &a, &b); a == 9999.9 && b == 9999.9 {
			break
		}
		count++
		testIn(count, point{a, b}, rectangles)
	}
}

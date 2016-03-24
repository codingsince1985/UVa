// UVa 477 - Points in Figures: Rectangles and Circles

package main

import (
	"fmt"
	"os"
)

type rectangle struct{ x1, y1, x2, y2 float64 }

type circle struct{ x, y, radius float64 }

type shape interface {
	contains(x, y float64) bool
}

func (r rectangle) contains(x, y float64) bool { return x > r.x1 && x < r.x2 && y < r.y1 && y > r.y2 }

func (c circle) contains(x, y float64) bool {
	return (x-c.x)*(x-c.x)+(y-c.y)*(y-c.y) < c.radius*c.radius
}

func testIn(out *os.File, count int, x, y float64, shapes []shape) {
	var inAny bool
	for i, v := range shapes {
		if v.contains(x, y) {
			fmt.Fprintf(out, "Point %d is contained in figure %d\n", count, i+1)
			inAny = true
		}
	}
	if !inAny {
		fmt.Fprintf(out, "Point %d is not contained in any figure\n", count)
	}
}

func main() {
	in, _ := os.Open("477.in")
	defer in.Close()
	out, _ := os.Create("477.out")
	defer out.Close()

	var shapes []shape
	var a, b, c, d float64
	var str string
	for {
		if fmt.Fscanf(in, "%s", &str); str == "*" {
			break
		}
		switch str {
		case "r":
			fmt.Fscanf(in, "%f%f%f%f", &a, &b, &c, &d)
			shapes = append(shapes, rectangle{a, b, c, d})
		case "c":
			fmt.Fscanf(in, "%f%f%f", &a, &b, &c)
			shapes = append(shapes, circle{a, b, c})
		}
	}

	count := 0
	for {
		if fmt.Fscanf(in, "%f%f", &a, &b); a == 9999.9 && b == 9999.9 {
			break
		}
		count++
		testIn(out, count, a, b, shapes)
	}
}

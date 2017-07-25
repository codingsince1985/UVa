// UVa 477 - Points in Figures: Rectangles and Circles

package main

import (
	"fmt"
	"os"
)

var out *os.File

type (
	point     struct{ x, y float64 }
	rectangle struct{ p1, p2 point }
	circle    struct {
		point
		radius float64
	}
	shape interface {
		contains(p point) bool
	}
)

func (r rectangle) contains(p point) bool {
	return p.x > r.p1.x && p.x < r.p2.x && p.y < r.p1.y && p.y > r.p2.y
}

func (c circle) contains(p point) bool {
	return (p.x-c.x)*(p.x-c.x)+(p.y-c.y)*(p.y-c.y) < c.radius*c.radius
}

func testIn(count int, p point, shapes []shape) {
	inAny := false
	for i, v := range shapes {
		if v.contains(p) {
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
	out, _ = os.Create("477.out")
	defer out.Close()

	var shapes []shape
	var c float64
	var p1, p2 point
	var str string
	for {
		if fmt.Fscanf(in, "%s", &str); str == "*" {
			break
		}
		switch str {
		case "r":
			fmt.Fscanf(in, "%f%f%f%f", &p1.x, &p1.y, &p2.x, &p2.y)
			shapes = append(shapes, rectangle{p1, p2})
		case "c":
			fmt.Fscanf(in, "%f%f%f", &p1.x, &p1.y, &c)
			shapes = append(shapes, circle{p1, c})
		}
	}

	for count := 1; ; count++ {
		if fmt.Fscanf(in, "%f%f", &p1.x, &p1.y); p1.x == 9999.9 && p1.y == 9999.9 {
			break
		}
		testIn(count, p1, shapes)
	}
}

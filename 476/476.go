// UVa 476 - Points in Figures: Rectangles

package main

import (
	"fmt"
	"io"
	"os"
)

type (
	point     struct{ x, y float64 }
	rectangle struct{ p1, p2 point }
)

func testIn(out io.Writer, count int, p point, rectangles []rectangle) {
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
	out, _ := os.Create("476.out")
	defer out.Close()

	var rectangles []rectangle
	var str string
	var p1, p2 point
	for {
		if fmt.Fscanf(in, "%s", &str); str == "*" {
			break
		}
		fmt.Fscanf(in, "%f%f%f%f", &p1.x, &p1.y, &p2.x, &p2.y)
		rectangles = append(rectangles, rectangle{p1, p2})
	}

	for count := 1; ; count++ {
		if fmt.Fscanf(in, "%f%f", &p1.x, &p1.y); p1.x == 9999.9 && p1.y == 9999.9 {
			break
		}
		testIn(out, count, p1, rectangles)
	}
}

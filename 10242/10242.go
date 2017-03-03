// UVa 10242 - Fourth Point !!

package main

import (
	"fmt"
	"os"
)

type (
	point struct{ x, y float64 }
	line  struct{ s, e point }
)

func connect(l1, l2 *line) {
	switch {
	case l1.s == l2.s:
		(*l1).s, (*l1).e = (*l1).e, (*l1).s
	case l1.e == l2.e:
		(*l2).s, (*l2).e = (*l2).e, (*l2).s
	case l1.s == l2.e:
		*l1, *l2 = *l2, *l1
	}
}

func main() {
	in, _ := os.Open("10242.in")
	defer in.Close()
	out, _ := os.Create("10242.out")
	defer out.Close()

	var l1, l2 line
	for {
		if _, err := fmt.Fscanf(in, "%f%f%f%f", &l1.s.x, &l1.s.y, &l1.e.x, &l1.e.y); err != nil {
			break
		}
		fmt.Fscanf(in, "%f%f%f%f", &l2.s.x, &l2.s.y, &l2.e.x, &l2.e.y)
		connect(&l1, &l2)
		fmt.Fprintf(out, "%.3f %.3f\n", l2.e.x-(l1.e.x-l1.s.x), l2.e.y-(l1.e.y-l1.s.y))
	}
}

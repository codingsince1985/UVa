// UVa 11498 - Division of Nlogonia

package main

import (
	"fmt"
	"os"
)

type point struct{ x, y int }

func country(division point, p point) string {
	switch {
	case p.x == division.x || p.y == division.y:
		return "divisa"
	case p.x > division.x:
		if p.y > division.y {
			return "NE"
		}
		return "SE"
	default:
		if p.y > division.y {
			return "NO"
		}
		return "SO"
	}
}

func main() {
	in, _ := os.Open("11498.in")
	defer in.Close()
	out, _ := os.Create("11498.out")
	defer out.Close()

	var k int
	var division, p point
	for {
		if fmt.Fscanf(in, "%d", &k); k == 0 {
			break
		}
		fmt.Fscanf(in, "%d%d", &division.x, &division.y)
		for ; k > 0; k-- {
			fmt.Fscanf(in, "%d%d", &p.x, &p.y)
			fmt.Fprintln(out, country(division, p))
		}
	}
}

// UVa 587 - There's treasure everywhere!

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const diagonal = math.Sqrt2 / 2

var directions = map[string][2]float64{
	"N": {0, 1}, "NE": {diagonal, diagonal},
	"E": {1, 0}, "SE": {diagonal, -diagonal},
	"S": {0, -1}, "SW": {-diagonal, -diagonal},
	"W": {-1, 0}, "NW": {-diagonal, diagonal},
}

func follow(steps []string) (float64, float64) {
	var x, y float64
	var length int
	var direction string
	for _, step := range steps {
		fmt.Sscanf(step, "%d%s", &length, &direction)
		vector := directions[direction]
		x += float64(length) * vector[0]
		y += float64(length) * vector[1]
	}
	return x, y
}

func main() {
	in, _ := os.Open("587.in")
	defer in.Close()
	out, _ := os.Create("587.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase int
	for s.Scan() {
		if line := s.Text(); line != "END" {
			x, y := follow(strings.Split(line[:len(line)-1], ","))
			kase++
			fmt.Fprintf(out, "Map #%d\n", kase)
			fmt.Fprintf(out, "The treasure is located at (%.3f,%.3f).\n", x, y)
			fmt.Fprintf(out, "The distance to the treasure is %.3f.\n\n", math.Sqrt(x*x+y*y))
		}
	}
}

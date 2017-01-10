// UVa 10573 - Geometry Paradox

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	in, _ := os.Open("10573.in")
	defer in.Close()
	out, _ := os.Create("10573.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var r1, r2 float64
	s.Scan()
	for kase, _ := strconv.Atoi(s.Text()); kase > 0 && s.Scan(); kase-- {
		if n, _ := fmt.Sscanf(s.Text(), "%f%f", &r1, &r2); n == 1 {
			fmt.Fprintf(out, "%.4f\n", math.Pi*r1*r1/8)
		} else {
			fmt.Fprintf(out, "%.4f\n", math.Pi*r1*r2*2)
		}
	}
}

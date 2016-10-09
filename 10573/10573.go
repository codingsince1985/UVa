// UVa 10573 - Geometry Paradox

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	in, _ := os.Open("10573.in")
	defer in.Close()
	out, _ := os.Create("10573.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	var r1, r2 float64
	for kase > 0 && s.Scan() {
		r := strings.NewReader(s.Text())
		if n, _ := fmt.Fscanf(r, "%f%f", &r1, &r2); n == 1 {
			fmt.Fprintf(out, "%.4f\n", math.Pi*r1*r1/8)
		} else {
			fmt.Fprintf(out, "%.4f\n", math.Pi*r1*r2*2)
		}
		kase--
	}
}

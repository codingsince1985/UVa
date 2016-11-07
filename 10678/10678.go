// UVa 10678 - The Grazing Cow

package main

import (
	"fmt"
	"math"
	"os"
)

func area(d, l int) float64 {
	var hl = float64(l) / 2
	var hd = float64(d) / 2
	a := math.Sqrt(hl*hl - hd*hd)
	return math.Pi * a * hl
}

func main() {
	in, _ := os.Open("10678.in")
	defer in.Close()
	out, _ := os.Create("10678.out")
	defer out.Close()

	var n, d, l int
	fmt.Fscanf(in, "%d", &n)
	for n > 0 {
		fmt.Fscanf(in, "%d%d", &d, &l)
		fmt.Fprintf(out, "%.3f\n", area(d, l))
		n--
	}
}

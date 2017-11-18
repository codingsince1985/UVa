// UVa 545 - Heads

package main

import (
	"fmt"
	"math"
	"os"
)

var log2 = math.Log10(2)

func solve(n int) (float64, int) {
	lg := float64(-n) * log2
	e := math.Floor(lg)
	return math.Pow(10, lg-e), int(e)
}

func main() {
	in, _ := os.Open("545.in")
	defer in.Close()
	out, _ := os.Create("545.out")
	defer out.Close()

	var r, n int
	for fmt.Fscanf(in, "%d", &r); r > 0; r-- {
		fmt.Fscanf(in, "%d", &n)
		n1, n2 := solve(n)
		fmt.Fprintf(out, "2^-%d = %.3fE%d\n", n, n1, n2)
	}
}

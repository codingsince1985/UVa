// UVa 10451 - Ancient Village Sports

package main

import (
	"fmt"
	"math"
	"os"
)

func calc(n, a float64) (float64, float64) {
	rs1 := a / (n * math.Tan(math.Pi/n))
	rs2 := 2 * a / (n * math.Sin(2*math.Pi/n))
	return math.Pi*rs2 - a, a - math.Pi*rs1
}

func main() {
	in, _ := os.Open("10451.in")
	defer in.Close()
	out, _ := os.Create("10451.out")
	defer out.Close()

	var kase int
	var n, a float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f", &n, &a); n < 3 || err != nil {
			break
		}
		kase++
		s, o := calc(n, a)
		fmt.Fprintf(out, "Case %d: %.5f %.5f\n", kase, s, o)
	}
}

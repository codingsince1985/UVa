// UVa 10522 - Height to Area

package main

import (
	"fmt"
	"math"
	"os"
)

func solve(ha, hb, hc float64) (bool, float64) {
	if ha == 0 || hb == 0 || hc == 0 {
		return false, 0
	}
	a, b, c := 1/ha, 1/hb, 1/hc
	p, pa, pb, pc := a+b+c, -a+b+c, a-b+c, a+b-c
	if s := p * pa * pb * pc; s > 0 {
		return true, 1 / math.Sqrt(s)
	}
	return false, 0
}

func main() {
	in, _ := os.Open("10522.in")
	defer in.Close()
	out, _ := os.Create("10522.out")
	defer out.Close()

	var invalid int
	var ha, hb, hc float64
	for fmt.Fscanf(in, "%d", &invalid); invalid > 0; {
		if _, err := fmt.Fscanf(in, "%f%f%f", &ha, &hb, &hc); err != nil {
			break
		}
		if ok, area := solve(ha, hb, hc); !ok {
			fmt.Fprintln(out, "These are invalid inputs!")
			invalid--
		} else {
			fmt.Fprintf(out, "%.3f\n", area)
		}
	}
}

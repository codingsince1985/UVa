// UVa 654 - Ratio

package main

import (
	"fmt"
	"io"
	"os"
)

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func solve(out io.Writer, gainers, losers int) {
	ratio, delta := float64(gainers)/float64(losers), 1.0
	for denominator := 1; denominator <= losers; denominator++ {
		numerator := int(float64(gainers*denominator)/float64(losers) + .5)
		if newRatio := abs(float64(numerator)/float64(denominator) - ratio); newRatio < delta {
			fmt.Fprintf(out, "%d/%d\n", numerator, denominator)
			delta = newRatio
		}
	}
}

func main() {
	in, _ := os.Open("654.in")
	defer in.Close()
	out, _ := os.Create("654.out")
	defer out.Close()

	var gainers, losers int
	for first := true; ; {
		if _, err := fmt.Fscanf(in, "%d%d", &gainers, &losers); err != nil {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		solve(out, gainers, losers)
	}
}

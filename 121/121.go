// UVa 121 - Pipe Fitters

package main

import (
	"fmt"
	"math"
	"os"
)

func skew(a, b float64) int {
	cols := int(a)
	rows := int(2*(b-1)/math.Sqrt(3)) + 1
	num := cols * rows
	if a-math.Floor(a) < 0.5 {
		num -= rows / 2
	}
	return num
}

func main() {
	in, _ := os.Open("121.in")
	defer in.Close()
	out, _ := os.Create("121.out")
	defer out.Close()

	var a, b float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f", &a, &b); err != nil {
			break
		}
		numG := int(a) * int(b)
		numS := max(skew(a, b), skew(b, a))
		maxNum := max(numG, numS)
		fmt.Fprintf(out, "%d ", maxNum)
		if maxNum == numG {
			fmt.Fprintln(out, "grid")
		} else {
			fmt.Fprintln(out, "skew")
		}
	}
}

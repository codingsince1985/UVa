// UVa 10347 - Medians

package main

import (
	"fmt"
	"math"
	"os"
)

func area(a, b, c float64) float64 {
	if a+b <= c || a+c <= b || b+c <= a {
		return -1
	}
	p := (a + b + c) / 2
	return math.Sqrt(p*(p-a)*(p-b)*(p-c)) * 4 / 3
}

func main() {
	in, _ := os.Open("10347.in")
	defer in.Close()
	out, _ := os.Create("10347.out")
	defer out.Close()

	var a, b, c float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f%f", &a, &b, &c); err != nil {
			break
		}
		fmt.Fprintf(out, "%.3f\n", area(a, b, c))
	}
}

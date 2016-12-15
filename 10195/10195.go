// UVa 10195 - The Knights Of The Round Table

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("10195.in")
	defer in.Close()
	out, _ := os.Create("10195.out")
	defer out.Close()

	var a, b, c, r float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f%f", &a, &b, &c); err != nil {
			break
		}
		if s := (a + b + c) / 2; s == 0 {
			r = 0
		} else {
			r = math.Sqrt(s*(s-a)*(s-b)*(s-c)) / s
		}
		fmt.Fprintf(out, "The radius of the round table is: %.3f\n", r)
	}
}

// UVa 10210 - Romeo and Juliet !

package main

import (
	"fmt"
	"math"
	"os"
)

type point struct{ x, y float64 }

func distance(p1, p2 point) float64 { return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)) }

func toRad(degree float64) float64 { return math.Pi * degree / 180 }

func main() {
	in, _ := os.Open("10210.in")
	defer in.Close()
	out, _ := os.Create("10210.out")
	defer out.Close()

	var a, b point
	var cmd, enf float64
	for {
		if _, err := fmt.Fscanf(in, "%f%f%f%f%f%f", &a.x, &a.y, &b.x, &b.y, &cmd, &enf); err != nil {
			break
		}
		fmt.Fprintf(out, "%.3f\n", distance(a, b)*(1/math.Tan(toRad(cmd))+1/math.Tan(toRad(enf))))
	}
}

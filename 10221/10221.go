// UVa 10221 - Satellites

package main

import (
	"fmt"
	"math"
	"os"
)

func degree2Radian(degree float64, unit string) float64 {
	if unit == "deg" {
		if degree > 180 {
			degree = 360 - degree
		}
		return degree * math.Pi / 180
	}
	return degree * math.Pi / 60 / 180
}

func main() {
	in, _ := os.Open("10221.in")
	defer in.Close()
	out, _ := os.Create("10221.out")
	defer out.Close()

	var s, a float64
	var unit string
	for {
		if _, err := fmt.Fscanf(in, "%f%f%s", &s, &a, &unit); err != nil {
			break
		}
		radius := 6440 + s
		radian := degree2Radian(a, unit)
		fmt.Fprintf(out, "%.6f %.6f\n", radius*radian, 2*radius*math.Sin(radian/2))
	}
}

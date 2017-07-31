// UVa 535 - Globetrotter

package main

import (
	"fmt"
	"math"
	"os"
)

const radius = 6378

type point struct{ lat, long float64 }

var cityMap = make(map[string]point)

func solve(p1, p2 point) float64 {
	dlat, dlong := p2.lat-p1.lat, p2.long-p1.long
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(p1.lat)*math.Cos(p2.lat)*math.Pow(math.Sin(dlong/2), 2)
	return radius * 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

func main() {
	in, _ := os.Open("535.in")
	defer in.Close()
	out, _ := os.Create("535.out")
	defer out.Close()

	var c1, c2 string
	var lat, long float64
	for {
		if fmt.Fscanf(in, "%s", &c1); c1 == "#" {
			break
		}
		fmt.Fscanf(in, "%f%f", &lat, &long)
		cityMap[c1] = point{lat * math.Pi / 180, long * math.Pi / 180}
	}
	for {
		if fmt.Fscanf(in, "%s%s", &c1, &c2); c1 == "#" && c2 == "#" {
			break
		}
		fmt.Fprintf(out, "%s - %s\n", c1, c2)
		if p1, ok := cityMap[c1]; ok {
			if p2, ok := cityMap[c2]; ok {
				fmt.Fprintf(out, "%.0f km\n", solve(p1, p2))
				continue
			}
		}
		fmt.Fprintln(out, "Unknown")
	}
}

// UVa 10060 - A hole to catch a man

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type (
	point struct{ x, y float64 }
	sheet struct {
		thickness float64
		points    []point
	}
)

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func solve(sheets []sheet, radius, thickness float64) int {
	var volume float64
	for _, s := range sheets {
		var area float64
		for i := 0; i < len(s.points)-1; i++ {
			area += s.points[i].x*s.points[i+1].y - s.points[i].y*s.points[i+1].x
		}
		volume += abs(area) / 2 * s.thickness
	}
	return int(volume / (math.Pi * radius * radius * thickness))
}

func main() {
	in, _ := os.Open("10060.in")
	defer in.Close()
	out, _ := os.Create("10060.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n int
	var tmp point
	var radius, thickness float64
	for s.Scan() {
		if fmt.Sscanf(s.Text(), "%d", &n); n == 0 {
			break
		}
		sheets := make([]sheet, n)
		for i := range sheets {
			s.Scan()
			r := strings.NewReader(s.Text())
			fmt.Fscanf(r, "%f", &sheets[i].thickness)
			for {
				if _, err := fmt.Fscanf(r, "%f%f", &tmp.x, &tmp.y); err != nil {
					break
				}
				sheets[i].points = append(sheets[i].points, tmp)
			}
		}
		s.Scan()
		fmt.Sscanf(s.Text(), "%f%f", &radius, &thickness)
		fmt.Fprintln(out, solve(sheets, radius, thickness))
	}
}

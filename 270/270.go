// UVa 270 - Lining Up

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type point struct{ x, y float64 }

func slope(p1, p2 point) [2]float64 {
	if p1.x == p2.x {
		return [2]float64{math.MaxFloat64, 0}
	}
	k := (p1.y - p2.y) / (p1.x - p2.x)
	return [2]float64{k, -k*p1.x + p1.y}
}

func solve(points []point) int {
	slopeMap := make(map[[2]float64]int)
	for i, l := 0, len(points); i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			slopeMap[slope(points[i], points[j])]++
		}
	}
	var max int
	for _, v := range slopeMap {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	in, _ := os.Open("270.in")
	defer in.Close()
	out, _ := os.Create("270.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	s.Scan()
	var x, y float64
	var line string
	for ; kase > 0; kase-- {
		var points []point
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			fmt.Sscanf(s.Text(), "%f%f", &x, &y)
			points = append(points, point{x, y})
		}
		fmt.Fprintln(out, solve(points))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}

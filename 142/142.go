// UVa 142 - Mouse Clicks

package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type (
	point  struct{ x, y int }
	region struct{ p1, p2 point }
)

var (
	icons   []point
	regions []region
)

func inRegion(p point, r region) bool {
	return p.x >= r.p1.x && p.x <= r.p2.x && p.y >= r.p1.y && p.y <= r.p2.y
}

func distance(p1, p2 point) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func solve(p point) string {
	for i := len(regions) - 1; i >= 0; i-- {
		if inRegion(p, regions[i]) {
			return string('A' + i)
		}
	}

	var nums []string
	closest := math.MaxFloat64
	for i, icon := range icons {
		dist := distance(p, icon)
		switch {
		case dist < closest:
			nums = nil
			closest = dist
			fallthrough
		case dist == closest:
			nums = append(nums, fmt.Sprintf("%3d", i+1))
		}
	}
	return strings.Join(nums, "")
}

func main() {
	in, _ := os.Open("142.in")
	defer in.Close()
	out, _ := os.Create("142.out")
	defer out.Close()

	var x1, y1, x2, y2 int
	var dataType byte
here:
	for {
		fmt.Fscanf(in, "%c", &dataType)
		switch dataType {
		case 'I':
			fmt.Fscanf(in, "%d%d", &x1, &y1)
			icons = append(icons, point{x1, y1})
		case 'R':
			fmt.Fscanf(in, "%d%d%d%d", &x1, &y1, &x2, &y2)
			regions = append(regions, region{point{x1, y1}, point{x2, y2}})
		case 'M':
			fmt.Fscanf(in, "%d%d", &x1, &y1)
			fmt.Fprintln(out, solve(point{x1, y1}))
		case '#':
			break here
		}
	}
}

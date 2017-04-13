// UVa 10369 - Arctic Network

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type (
	outpost struct{ x, y float64 }
	edge    struct {
		o1, o2   int
		distance float64
	}
	edges []edge
)

func (e edges) Len() int           { return len(e) }
func (e edges) Less(i, j int) bool { return e[i].distance < e[j].distance }
func (e edges) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func distance(o1, o2 outpost) float64 {
	return math.Sqrt((o1.x-o2.x)*(o1.x-o2.x) + (o1.y-o2.y)*(o1.y-o2.y))
}

func solve(s, p int, outposts []outpost) float64 {
	var e edges
	for i := 0; i < p-1; i++ {
		for j := i + 1; j < p; j++ {
			e = append(e, edge{i, j, distance(outposts[i], outposts[j])})
		}
	}
	sort.Sort(e)
	var maxDistance float64
	var idx int
	f := make([]int, p)
	for i := range f {
		f[i] = i
	}
	for _, ve := range e {
		if f1, f2 := unionFind(ve.o1, f), unionFind(ve.o2, f); f1 != f2 {
			f[f1] = f2
			maxDistance = ve.distance
			if idx++; idx == p-s {
				break
			}
		}
	}
	return maxDistance
}

func main() {
	in, _ := os.Open("10369.in")
	defer in.Close()
	out, _ := os.Create("10369.out")
	defer out.Close()

	var n, s, p int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d%d", &s, &p)
		outposts := make([]outpost, p)
		for i := range outposts {
			fmt.Fscanf(in, "%f%f", &outposts[i].x, &outposts[i].y)
		}
		fmt.Fprintf(out, "%.2f\n", solve(s, p, outposts))
	}
}

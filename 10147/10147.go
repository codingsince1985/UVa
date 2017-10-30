// UVa 10147 - Highways

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type (
	town    struct{ x, y float64 }
	highway struct {
		distance float64
		t1, t2   int
	}
)

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func distance(t1, t2 town) float64 {
	return math.Sqrt((t1.x-t2.x)*(t1.x-t2.x) + (t1.y-t2.y)*(t1.y-t2.y))
}

func solve(out *os.File, n int, towns []town, f []int) {
	highways := buildHighways(n, towns)
	noNew := true
	for _, h := range highways {
		if fa, fb := unionFind(h.t1, f), unionFind(h.t2, f); fa != fb {
			noNew = false
			f[fa] = fb
			fmt.Fprintln(out, h.t1+1, h.t2+1)
		}
	}
	if noNew {
		fmt.Fprintln(out, "‘No new highways need")
	}
}

func buildHighways(n int, towns []town) []highway {
	var highways []highway
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			highways = append(highways, highway{distance(towns[i], towns[j]), i, j})
		}
	}
	sort.Slice(highways, func(i, j int) bool { return highways[i].distance < highways[j].distance })
	return highways
}

func main() {
	in, _ := os.Open("10147.in")
	defer in.Close()
	out, _ := os.Create("10147.out")
	defer out.Close()

	var kase, n, m, t1, t2 int
	first := true
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &n)
		towns := make([]town, n)
		for i := range towns {
			fmt.Fscanf(in, "%f%f", &towns[i].x, &towns[i].y)
		}
		f := make([]int, n)
		for i := range f {
			f[i] = i
		}
		for fmt.Fscanf(in, "%d", &m); m > 0; m-- {
			fmt.Fscanf(in, "%d%d", &t1, &t2)
			if fa, fb := unionFind(t1-1, f), unionFind(t2-1, f); fa != fb {
				f[fa] = fb
			}
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		solve(out, n, towns, f)
	}
}

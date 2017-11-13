// UVa 10022 - Delta-wave

package main

import (
	"fmt"
	"math"
	"os"
)

type point struct{ x, y int }

func locate(n int) point {
	for h := 1; ; h++ {
		if h*h >= n {
			return point{h, n - (h-1)*(h-1)}
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func solve(from, to point) int {
	if from.x == to.x {
		return abs(from.y - to.y)
	}
	steps := math.MaxInt32
	if from.y%2 == 1 {
		for x, y := to.x, 1; y <= 2*x-1; y++ {
			if from.y <= y && y <= from.y+2*(to.x-from.x) {
				if tmp := 2*(x-from.x) + abs(y-to.y); y%2 == 1 {
					steps = min(steps, tmp)
				} else {
					steps = min(steps, tmp-1)
				}
			}
		}
	} else {
		for x, y := from.x, 1; y <= 2*x-1; y += 2 {
			if y <= to.y && to.y <= y+2*(to.x-from.x) {
				if tmp := 2*(to.x-x) + abs(y-from.y); to.y%2 == 1 {
					steps = min(steps, tmp)
				} else {
					steps = min(steps, tmp-1)
				}
			}
		}
	}
	return steps
}

func main() {
	in, _ := os.Open("10022.in")
	defer in.Close()
	out, _ := os.Create("10022.out")
	defer out.Close()

	var kase, m, n int
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanln(in)
		fmt.Fscanf(in, "%d%d", &m, &n)
		fmt.Fprintln(out, solve(locate(m), locate(n)))
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}

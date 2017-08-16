// UVa 10088 - Trees on My Island

package main

import (
	"fmt"
	"os"
)

type point struct{ x, y int64 }

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func solve(n int, vertices []point) int64 {
	var area int64
	for i := 0; i < n; i++ {
		area += (vertices[i].x * vertices[i+1].y) - (vertices[i].y * vertices[i+1].x)
	}
	var b int64
	for i := 0; i < n; i++ {
		b += gcd(abs(vertices[i].x-vertices[i+1].x), abs(vertices[i].y-vertices[i+1].y))
	}
	return (abs(area) + 2 - b) / 2
}

func main() {
	in, _ := os.Open("10088.in")
	defer in.Close()
	out, _ := os.Create("10088.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		vertices := make([]point, n)
		for i := range vertices {
			fmt.Fscanf(in, "%d%d", &vertices[i].x, &vertices[i].y)
		}
		fmt.Fprintln(out, solve(n, append(vertices, point{vertices[0].x, vertices[0].y})))
	}
}

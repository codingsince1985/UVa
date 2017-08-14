// UVa 273 - Jack Straws

package main

import (
	"fmt"
	"os"
)

type (
	point struct{ x, y int }
	line  struct{ p1, p2 point }
)

func area(p point, l line) int { return (l.p1.y-p.y)*(l.p2.x-p.x) - (l.p2.y-p.y)*(l.p1.x-p.x) }

func cross(p point, l line) bool {
	return (p.x-l.p1.x)*(p.x-l.p2.x) <= 0 && (p.y-l.p1.y)*(p.y-l.p2.y) <= 0
}

func intersect(l1, l2 line) bool {
	a1 := area(l2.p1, l1)
	a2 := area(l2.p2, l1)
	a3 := area(l1.p1, l2)
	a4 := area(l1.p2, l2)

	return a1*a2 < 0 && a3*a4 < 0 ||
		a1 == 0 && cross(l2.p1, l1) ||
		a2 == 0 && cross(l2.p2, l1) ||
		a3 == 0 && cross(l1.p1, l2) ||
		a4 == 0 && cross(l1.p2, l2)
}

func floydWarshall(matrix [][]bool) {
	for k := range matrix {
		for i := range matrix {
			for j := range matrix {
				if i != j && matrix[i][k] && matrix[k][j] && !matrix[i][j] {
					matrix[i][j], matrix[j][i] = true, true
				}
			}
		}
	}
}

func solve(n int, straws []line) [][]bool {
	connected := make([][]bool, n)
	for i := range connected {
		connected[i] = make([]bool, n)
		connected[i][i] = true
	}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if intersect(straws[i], straws[j]) {
				connected[i][j], connected[j][i] = true, true
			}
		}
	}
	floydWarshall(connected)
	return connected
}

func main() {
	in, _ := os.Open("273.in")
	defer in.Close()
	out, _ := os.Create("273.out")
	defer out.Close()

	var kase, n, s1, s2 int
	first := true
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "\n%d", &n)
		straws := make([]line, n)
		for i := range straws {
			fmt.Fscanf(in, "%d%d%d%d", &straws[i].p1.x, &straws[i].p1.y, &straws[i].p2.x, &straws[i].p2.y)
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		connected := solve(n, straws)
		for {
			if fmt.Fscanf(in, "%d%d", &s1, &s2); s1 == 0 && s2 == 0 {
				break
			}
			if connected[s1-1][s2-1] {
				fmt.Fprintln(out, "CONNECTED")
			} else {
				fmt.Fprintln(out, "NOT CONNECTED")
			}
		}
	}
}

// UVa 10112 - Myacm Triangles

package main

import (
	"fmt"
	"os"
)

const zero = 0.00001

type (
	monument struct {
		label byte
		x, y  float64
	}
	triangle struct{ m1, m2, m3 monument }
)

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func area(m1, m2, m3 monument) float64 {
	return abs(0.5 * ((m3.y-m1.y)*(m2.x-m1.x) - (m2.y-m1.y)*(m3.x-m1.x)))
}

func (t triangle) contains(p monument) bool {
	return abs(area(p, t.m1, t.m2)+area(p, t.m1, t.m3)+area(p, t.m2, t.m3)-area(t.m1, t.m2, t.m3)) <= zero
}

func solve(n int, monuments []monument) string {
	var labels []byte
	var max float64
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
		here:
			for k := j + 1; k < n; k++ {
				if a := area(monuments[i], monuments[j], monuments[k]); a > max {
					for l := 0; l < n; l++ {
						if l != i && l != j && l != k {
							t := triangle{monuments[i], monuments[j], monuments[k]}
							if t.contains(monuments[l]) {
								continue here
							}
						}
					}
					max = a
					labels = []byte{monuments[i].label, monuments[j].label, monuments[k].label}
				}
			}
		}
	}
	return string(labels)
}

func main() {
	in, _ := os.Open("10112.in")
	defer in.Close()
	out, _ := os.Create("10112.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		monuments := make([]monument, n)
		for i := range monuments {
			fmt.Fscanf(in, "%c%f%f", &monuments[i].label, &monuments[i].x, &monuments[i].y)
		}
		fmt.Fprintln(out, solve(n, monuments))
	}
}

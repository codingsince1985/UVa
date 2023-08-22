// UVa 10354 - Avoiding Your Boss

package main

import (
	"fmt"
	"math"
	"os"
)

func floydWarshall(places [][]int) {
	for k := range places {
		for i := range places {
			for j := range places {
				places[i][j] = min(places[i][j], places[i][k]+places[k][j])
			}
		}
	}
}

func floydWarshall2(places [][]int, avoided []bool) {
	for k := range places {
		if !avoided[k] {
			for i := range places {
				if !avoided[i] {
					for j := range places {
						if !avoided[j] {
							places[i][j] = min(places[i][j], places[i][k]+places[k][j])
						}
					}
				}
			}
		}
	}
}

func solve(p, bh, of, yh, m int, places [][]int) (bool, int) {
	distances := make([][]int, p)
	for i := range distances {
		distances[i] = make([]int, p)
	}
	copy(distances, places)
	floydWarshall(places)
	avoided := make([]bool, p)
	for i := range places {
		if places[bh][of] == places[bh][i]+places[i][of] {
			avoided[i] = true
		}
	}
	if avoided[yh] || avoided[m] {
		return false, -1
	}
	floydWarshall2(distances, avoided)
	if distances[yh][m] == math.MaxInt32 {
		return false, -1
	}
	return true, distances[yh][m]
}

func main() {
	in, _ := os.Open("10354.in")
	defer in.Close()
	out, _ := os.Create("10354.out")
	defer out.Close()

	var p, r, bh, of, yh, m, p1, p2, d int
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d%d%d%d", &p, &r, &bh, &of, &yh, &m); err != nil {
			break
		}
		places := make([][]int, p)
		for i := range places {
			places[i] = make([]int, p)
			for j := range places[i] {
				if i != j {
					places[i][j] = math.MaxInt32
				}
			}
		}
		for ; r > 0; r-- {
			fmt.Fscanf(in, "%d%d%d", &p1, &p2, &d)
			places[p1-1][p2-1], places[p2-1][p1-1] = d, d
		}
		if ok, distance := solve(p, bh-1, of-1, yh-1, m-1, places); ok {
			fmt.Fprintln(out, distance)
		} else {
			fmt.Fprintln(out, "MISSION IMPOSSIBLE.")
		}
	}
}

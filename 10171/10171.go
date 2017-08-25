// UVa 10171 - Meeting Prof. Miguel...

package main

import (
	"fmt"
	"math"
	"os"
)

func floydWarshall(matrix [][]int) {
	for k := range matrix {
		for i := range matrix {
			for j := range matrix {
				if energy := matrix[i][k] + matrix[k][j]; energy < matrix[i][j] {
					matrix[i][j] = energy
				}
			}
		}
	}
}

func solve(p1, p2 byte, young, old [][]int, placeMap map[byte]int) (int, byte) {
	floydWarshall(young)
	floydWarshall(old)
	least := math.MaxInt32
	var p int
	for i := range young {
		if energy := young[placeMap[p1]][i] + old[placeMap[p2]][i]; energy < least {
			least = energy
			p = i
		}
	}
	if least == math.MaxInt32 {
		return least, 0
	}
	for k, v := range placeMap {
		if v == p {
			return least, k
		}
	}
	return 0, 0
}

func main() {
	in, _ := os.Open("10171.in")
	defer in.Close()
	out, _ := os.Create("10171.out")
	defer out.Close()

	var n int
	var p1, p2 byte
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		ym, ub, x, y := make([]byte, n), make([]byte, n), make([]byte, n), make([]byte, n)
		c := make([]int, n)
		placeMap := make(map[byte]int)
		index := 0
		for i := range c {
			fmt.Fscanf(in, "%c %c %c %c %d", &ym[i], &ub[i], &x[i], &y[i], &c[i])
			if _, ok := placeMap[x[i]]; !ok {
				placeMap[x[i]] = index
				index++
			}
			if _, ok := placeMap[y[i]]; !ok {
				placeMap[y[i]] = index
				index++
			}
		}
		young, old := make([][]int, len(placeMap)), make([][]int, len(placeMap))
		for i := range young {
			young[i], old[i] = make([]int, len(placeMap)), make([]int, len(placeMap))
			for j := range young[i] {
				if i == j {
					young[i][j], old[i][j] = 0, 0
				} else {
					young[i][j], old[i][j] = math.MaxInt32, math.MaxInt32
				}
			}
		}
		for i, energy := range c {
			if ym[i] == 'Y' {
				young[placeMap[x[i]]][placeMap[y[i]]] = energy
				if ub[i] == 'B' {
					young[placeMap[y[i]]][placeMap[x[i]]] = energy
				}
			} else {
				old[placeMap[x[i]]][placeMap[y[i]]] = energy
				if ub[i] == 'B' {
					old[placeMap[y[i]]][placeMap[x[i]]] = energy
				}
			}
		}
		fmt.Fscanf(in, "%c %c\n", &p1, &p2)
		if least, p := solve(p1, p2, young, old, placeMap); least == math.MaxInt32 {
			fmt.Fprintln(out, "You will never meet.")
		} else {
			fmt.Fprintf(out, "%d %c\n", least, p)
		}
	}
}

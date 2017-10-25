// UVa 102 - Ecological Bin Packing

package main

import (
	"fmt"
	"math"
	"os"
)

const (
	b = iota
	c
	g
)

var (
	binColor = [][3]int{{b, c, g}, {b, g, c}, {c, b, g}, {c, g, b}, {g, b, c}, {g, c, b}}
	binCode  = []byte{'B', 'C', 'G'}
)

func main() {
	in, _ := os.Open("102.in")
	defer in.Close()
	out, _ := os.Create("102.out")
	defer out.Close()

	var bin [3][3]int
	var idx int
	for {
		total := 0
		for i := 0; i < 3; i++ {
			if _, err := fmt.Fscanf(in, "%d%d%d", &bin[i][b], &bin[i][g], &bin[i][c]); err != nil {
				return
			}
			total += bin[i][b] + bin[i][g] + bin[i][c]
		}
		minMove := math.MaxInt32
		for i, v := range binColor {
			if move := total - bin[0][v[0]] - bin[1][v[1]] - bin[2][v[2]]; move < minMove {
				minMove = move
				idx = i
			}
		}
		fmt.Fprintf(out, "%c%c%c %d\n", binCode[binColor[idx][0]], binCode[binColor[idx][1]], binCode[binColor[idx][2]], minMove)
	}
}

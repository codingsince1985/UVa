// UVa 102 - Ecological Bin Packing

package main

import (
	"fmt"
	"math"
	"os"
)

const (
	B = iota
	C
	G
)

var (
	binColor = [][3]int{{B, C, G}, {B, G, C}, {C, B, G}, {C, G, B}, {G, B, C}, {G, C, B}}
	binCode  = []byte{'B', 'C', 'G'}
)

func main() {
	in, _ := os.Open("102.in")
	defer in.Close()
	out, _ := os.Create("102.out")
	defer out.Close()

	var bin [3][3]int
	var total, minMove, idx int
	for {
		total = 0
		for i := 0; i < 3; i++ {
			if _, err := fmt.Fscanf(in, "%d%d%d", &bin[i][B], &bin[i][G], &bin[i][C]); err != nil {
				return
			}
			total += bin[i][B] + bin[i][G] + bin[i][C]
		}
		minMove = math.MaxInt32
		for i, v := range binColor {
			if move := total - bin[0][v[0]] - bin[1][v[1]] - bin[2][v[2]]; move < minMove {
				minMove = move
				idx = i
			}
		}
		fmt.Fprintf(out, "%c%c%c %d\n", binCode[binColor[idx][0]], binCode[binColor[idx][1]], binCode[binColor[idx][2]], minMove)
	}
}

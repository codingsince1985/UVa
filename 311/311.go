// UVa 311 - Packets

package main

import (
	"fmt"
	"os"
)

func pack(p [7]int) int {
	parcels := p[6] + p[5] + p[4]
	p[1] -= 11 * p[5]
	p[2] -= 5 * p[4]

	parcels += p[3] / 4
	switch p[3] % 4 {
	case 1:
		p[2] -= 5
		p[1] -= 7
		parcels++
	case 2:
		p[2] -= 3
		p[1] -= 6
		parcels++
	case 3:
		p[2]--
		p[1] -= 5
		parcels++
	}

	if p[2] > 0 {
		parcels += p[2] / 9
		if p[2]%9 != 0 {
			parcels++
		}
		p[1] -= 36 - p[2]%9*4
	} else {
		p[1] += p[2] * 4
	}

	if p[1] > 0 {
		parcels += p[1] / 36
		if p[1]%36 != 0 {
			parcels++
		}
	}
	return parcels
}

func main() {
	in, _ := os.Open("311.in")
	defer in.Close()
	out, _ := os.Create("311.out")
	defer out.Close()

	var p [7]int
	for {
		fmt.Fscanf(in, "%d%d%d%d%d%d", &p[1], &p[2], &p[3], &p[4], &p[5], &p[6])
		if p[1]+p[2]+p[3]+p[4]+p[5]+p[6] == 0 {
			break
		}
		fmt.Fprintln(out, pack(p))
	}
}

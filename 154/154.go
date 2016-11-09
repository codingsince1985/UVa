// UVa 154 - Recycling

package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const TYPES = 5

func solve(city []map[byte]byte) int {
	var idx, diff int
	min := math.MaxInt32
	for i, vi := range city {
		diff = 0
		for j, vj := range city {
			if i != j {
				for k := range vi {
					if vi[k] != vj[k] {
						diff++
					}
				}
			}
		}
		if diff < min {
			min, idx = diff, i
		}
	}
	return idx + 1
}

func main() {
	in, _ := os.Open("154.in")
	defer in.Close()
	out, _ := os.Create("154.out")
	defer out.Close()

	var line string
	var c1, c2, tmp byte
	var city []map[byte]byte
	for {
		if fmt.Fscanf(in, "%s", &line); line == "#" {
			break
		}
		if line[0] == 'e' {
			fmt.Fprintln(out, solve(city))
			city = nil
		} else {
			bin := make(map[byte]byte)
			r := strings.NewReader(line)
			for i := 0; i < TYPES; i++ {
				fmt.Fscanf(r, "%c/%c%c", &c1, &c2, &tmp)
				bin[c1] = c2
			}
			city = append(city, bin)
		}
	}
}

// UVa 10916 - Factstone Benchmark

package main

import (
	"fmt"
	"math"
	"os"
)

func solve(y uint) int {
	bits := 1 << ((y-1960)/10 + 2)
	var i int
	for sum := 0.0; sum < float64(bits); {
		i++
		sum += math.Log2(float64(i))
	}
	i--
	return i
}

func main() {
	in, _ := os.Open("10916.in")
	defer in.Close()
	out, _ := os.Create("10916.out")
	defer out.Close()

	var y uint
	for {
		if fmt.Fscanf(in, "%d", &y); y == 0 {
			break
		}
		fmt.Fprintln(out, solve(y))
	}
}

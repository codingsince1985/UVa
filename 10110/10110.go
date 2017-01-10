// UVa 10110 - Light, more light

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("10110.in")
	defer in.Close()
	out, _ := os.Create("10110.out")
	defer out.Close()

	var n uint64
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		if sr := uint64(math.Sqrt(float64(n))); sr*sr == n {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}
	}
}

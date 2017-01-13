// UVa 107 - The Cat in the Hat

package main

import (
	"fmt"
	"math"
	"os"
)

const diff = 1e-9

func main() {
	in, _ := os.Open("107.in")
	defer in.Close()
	out, _ := os.Create("107.out")
	defer out.Close()

	var h, c, n float64
	for {
		if fmt.Fscanf(in, "%f%f", &h, &c); h == 0 && c == 0 {
			break
		}
		for n = 0; math.Abs(math.Log(h)*math.Log(n)-math.Log(c)*math.Log(n+1)) > diff; n++ {
		}
		if n != 1 {
			fmt.Fprintf(out, "%.0f %.0f\n", (1-c)/(1-n), (h-c)*(n+1)+c)
		} else {
			fmt.Fprintf(out, "%.0f %.0f\n", math.Log2(h), (h-c)*(n+1)+c)
		}
	}
}

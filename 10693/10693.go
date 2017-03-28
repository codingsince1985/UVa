// UVa 10693 - Traffic Volume

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("10693.in")
	defer in.Close()
	out, _ := os.Create("10693.out")
	defer out.Close()

	var l, f float64
	for {
		if fmt.Fscanf(in, "%f%f", &l, &f); l == 0 && f == 0 {
			break
		}
		v := math.Sqrt(2 * l * f)
		fmt.Fprintf(out, "%.8f %.8f\n", v, 3600*v/(l+v*v/(2*f)))
	}
}

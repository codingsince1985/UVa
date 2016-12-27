// UVa 11636 - Hello World!

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("11636.in")
	defer in.Close()
	out, _ := os.Create("11636.out")
	defer out.Close()

	var n int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n < 0 {
			break
		}
		fmt.Fprintf(out, "Case %d: %.0f\n", kase, math.Ceil(math.Log2(float64(n))))
	}
}

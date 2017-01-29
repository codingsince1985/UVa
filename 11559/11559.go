// UVa 11559 - Event Planning

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	in, _ := os.Open("11559.in")
	defer in.Close()
	out, _ := os.Create("11559.out")
	defer out.Close()

	var n, b, h, w, p, a, total int
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d%d", &n, &b, &h, &w); err != nil {
			break
		}
		for total = math.MaxInt32; h > 0; h-- {
			fmt.Fscanf(in, "%d", &p)
			for i := 0; i < w; i++ {
				if fmt.Fscanf(in, "%d", &a); a >= n {
					if t := p * n; t <= b && t < total {
						total = t
					}
				}
			}
		}
		if total != math.MaxInt32 {
			fmt.Fprintln(out, total)
		} else {
			fmt.Fprintln(out, "stay home")
		}
	}
}

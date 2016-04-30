// UVa 573 - The Snail

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("573.in")
	defer in.Close()
	out, _ := os.Create("573.out")
	defer out.Close()

	var h, u, d, f float64
	for {
		if fmt.Fscanf(in, "%f%f%f%f", &h, &u, &d, &f); h == 0 {
			break
		}
		fatigue := u * f / 100
		pos := 0.0
		day := 0
		for {
			day++
			up := u - fatigue*float64(day-1)
			if up <= 0 {
				up = 0
			}
			pos += up
			if pos > h {
				fmt.Fprintf(out, "success on day %d\n", day)
				break
			}
			pos -= d
			if pos < 0 {
				fmt.Fprintf(out, "failure on day %d\n", day)
				break
			}
		}
	}
}

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
		for fatigue, pos, day := u*f/100, 0.0, 1; ; day++ {
			up := u - fatigue*float64(day-1)
			if up <= 0 {
				up = 0
			}
			if pos += up; pos > h {
				fmt.Fprintf(out, "success on day %d\n", day)
				break
			}
			if pos -= d; pos < 0 {
				fmt.Fprintf(out, "failure on day %d\n", day)
				break
			}
		}
	}
}

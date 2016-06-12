// UVa 10550 - Combination Lock

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10550.in")
	defer in.Close()
	out, _ := os.Create("10550.out")
	defer out.Close()

	var i0, i1, i2, i3 int
	for {
		if fmt.Fscanf(in, "%d%d%d%d", &i0, &i1, &i2, &i3); i0 == 0 && i1 == 0 && i2 == 0 && i3 == 0 {
			break
		}
		degree := 720
		if i0 > i1 {
			degree += (i0 - i1) * 9
		} else {
			degree += (i0 + 40 - i1) * 9
		}
		degree += 360
		if i2 > i1 {
			degree += (i2 - i1) * 9
		} else {
			degree += (i2 + 40 - i1) * 9
		}
		if i2 > i3 {
			degree += (i2 - i3) * 9
		} else {
			degree += (i2 + 40 - i3) * 9
		}
		fmt.Fprintln(out, degree)
	}
}

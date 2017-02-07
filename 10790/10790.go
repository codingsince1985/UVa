// UVa 10790 - How Many Points of Intersection?

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("10790.in")
	defer in.Close()
	out, _ := os.Create("10790.out")
	defer out.Close()

	var a, b int64
	for {
		if fmt.Fscanf(in, "%d%d", &a, &b); a == 0 && b == 0 {
			break
		}
		fmt.Fprintln(out, a*(a-1)*b*(b-1)/4)
	}
}

// UVa 11849 - CD

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("11849.in")
	defer in.Close()
	out, _ := os.Create("11849.out")
	defer out.Close()

	var n, m, catalog int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		jacks := make(map[int]bool)
		for ; n > 0; n-- {
			fmt.Fscanf(in, "%d", &catalog)
			jacks[catalog] = true
		}
		var count int
		for ; m > 0; m-- {
			if fmt.Fscanf(in, "%d", &catalog); jacks[catalog] {
				count++
			}
		}
		fmt.Fprintln(out, count)
	}
}

// UVa 320 - Border

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("320.in")
	defer in.Close()
	out, _ := os.Create("320.out")
	defer out.Close()

	var kase, r, c int
	var ch byte
	fmt.Fscanf(in, "%d", &kase)
	for i := 1; i <= kase; i++ {
		var grid [32][32]bool
		fmt.Fscanf(in, "%d%d", &c, &r)
		for {
			if fmt.Fscanf(in, "%c", &ch); ch == '.' {
				break
			}
			switch ch {
			case 'E':
				grid[r-1][c] = true
				c++
			case 'N':
				grid[r][c] = true
				r++
			case 'W':
				c--
				grid[r][c] = true
			case 'S':
				r--
				grid[r][c-1] = true
			}
		}
		fmt.Fprintf(out, "Bitmap #%d\n", i)
		for i := range grid {
			for j := range grid[31-i] {
				if grid[31-i][j] {
					ch = 'X'
				} else {
					ch = '.'
				}
				fmt.Fprintf(out, "%c", ch)
			}
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out)
	}
}

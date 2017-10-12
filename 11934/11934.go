// UVa 11934 - Magic Formula

package main

import (
	"fmt"
	"os"
)

func solve(a, b, c, d, l int64) int {
	var count int
	for x := int64(0); x <= l; x++ {
		if (a*x*x+b*x+c)%d == 0 {
			count++
		}
	}
	return count
}

func main() {
	in, _ := os.Open("11934.in")
	defer in.Close()
	out, _ := os.Create("11934.out")
	defer out.Close()

	var a, b, c, d, l int64
	for {
		if fmt.Fscanf(in, "%d%d%d%d%d", &a, &b, &c, &d, &l); a == 0 && b == 0 && c == 0 && d == 0 && l == 0 {
			break
		}
		fmt.Fprintln(out, solve(a, b, c, d, l))
	}
}

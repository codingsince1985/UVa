// UVa 11877 - The Coco-Cola Store

package main

import (
	"fmt"
	"os"
)

func solve(empty int) int {
	total := 0
	for empty >= 3 {
		toChange := empty / 3
		total += toChange
		empty -= (toChange * 3)
		empty += toChange
	}
	if empty == 2 {
		total++
	}
	return total
}

func main() {
	in, _ := os.Open("11877.in")
	defer in.Close()
	out, _ := os.Create("11877.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fprintln(out, solve(n))
	}
}
